/// <reference types="node" />
import { createServerAdapter as createAdapter } from '@whatwg-node/server';
import { type GraphQLSchema } from 'graphql';
import { createYoga } from 'graphql-yoga';
import type { IncomingMessage, ServerResponse } from 'node:http';
import type { RouterManifest, RouterPageManifest, YogaServerOptions } from './types';
export declare const serverAdapterFactory: <ComponentType>({ schema, yoga, production, manifest, on_render, pipe, assetPrefix, }: {
    schema?: GraphQLSchema | null | undefined;
    yoga?: import("graphql-yoga").YogaServerInstance<Record<string, any>, Record<string, any>> | null | undefined;
    assetPrefix: string;
    production?: boolean | undefined;
    pipe?: ServerResponse<IncomingMessage> | undefined;
    on_render: (args: {
        url: string;
        match: RouterPageManifest<ComponentType> | null;
        manifest: RouterManifest<unknown>;
        session: App.Session;
        pipe?: ServerResponse<IncomingMessage> | undefined;
    }) => Response | Promise<Response>;
    manifest: RouterManifest<ComponentType> | null;
} & Omit<import("graphql-yoga").YogaServerOptions<Record<string, any>, Record<string, any>>, "schema">) => ReturnType<typeof createAdapter>;
export type ServerAdapterFactory = typeof serverAdapterFactory;
