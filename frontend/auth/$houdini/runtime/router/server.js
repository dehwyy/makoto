import { parse } from "./cookies";
import { decode, encode, verify } from "./jwt";
async function handle_request(args) {
  const plugin_config = args.config.router ?? {};
  if (plugin_config.auth && "redirect" in plugin_config.auth && args.url.startsWith(plugin_config.auth.redirect)) {
    return await redirect_auth(args);
  }
  args.next();
}
async function redirect_auth(args) {
  const { searchParams } = new URL(args.url, `http://${args.get_header("host")}`);
  const { redirectTo, ...session } = Object.fromEntries(searchParams.entries());
  await set_session(args, session);
  if (redirectTo) {
    return args.redirect(302, redirectTo);
  }
  args.next();
}
const session_cookie_name = "__houdini__";
async function set_session(req, value) {
  const today = new Date();
  const expires = new Date(today.getTime() + 7 * 24 * 60 * 60 * 1e3);
  const serialized = await encode(value, req.session_keys[0]);
  req.set_header(
    "Set-Cookie",
    `${session_cookie_name}=${serialized}; Path=/; HttpOnly; Secure; SameSite=Lax; Expires=${expires.toUTCString()} `
  );
}
async function get_session(req, secrets) {
  const cookies = req.get("cookie");
  if (!cookies) {
    return {};
  }
  const cookie = parse(cookies)[session_cookie_name];
  if (!cookie) {
    return {};
  }
  for (const secret of secrets) {
    if (!await verify(cookie, secret)) {
      continue;
    }
    const parsed = decode(cookie);
    if (!parsed) {
      return {};
    }
    return parsed.payload;
  }
  return {};
}
export {
  get_session,
  handle_request
};
