import { createServerAdapter as createAdapter } from "@whatwg-node/server";
import { parse, execute } from "graphql";
import { createYoga } from "graphql-yoga";
import client from "../../../src/+client";
import { localApiSessionKeys, localApiEndpoint, getCurrentConfig } from "../lib/config";
import { find_match } from "./match";
import { get_session, handle_request } from "./session";
const config_file = getCurrentConfig();
const session_keys = localApiSessionKeys(config_file);
const graphqlEndpoint = localApiEndpoint(config_file);
const serverAdapterFactory = ({
  schema,
  yoga,
  production,
  manifest,
  on_render,
  pipe,
  assetPrefix
}) => {
  if (schema && !yoga) {
    yoga = createYoga({
      schema,
      landingPage: !production,
      graphqlEndpoint
    });
  }
  if (schema) {
    client.registerProxy(graphqlEndpoint, async ({ query, variables, session }) => {
      const parsed = parse(query);
      return await execute(schema, parsed, null, session, variables);
    });
  }
  return createAdapter(async (request) => {
    if (!manifest) {
      return new Response(
        "Adapter did not provide the project's manifest. Please open an issue on github.",
        { status: 500 }
      );
    }
    const url = new URL(request.url).pathname;
    if (yoga && url === localApiEndpoint(config_file)) {
      return yoga(request);
    }
    const authResponse = await handle_request({
      url,
      config: config_file,
      session_keys,
      headers: request.headers
    });
    if (authResponse) {
      return authResponse;
    }
    const [match] = find_match(manifest, url);
    const rendered = await on_render({
      url,
      match,
      session: await get_session(request.headers, session_keys),
      manifest,
      pipe
    });
    if (rendered) {
      console.log(url, rendered);
      return rendered;
    }
    return new Response("404", { status: 404 });
  });
};
export {
  serverAdapterFactory
};
