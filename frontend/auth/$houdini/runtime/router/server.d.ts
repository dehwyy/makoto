import type { ConfigFile } from '../lib';
type ServerHandlerArgs = {
    url: string;
    config: ConfigFile;
    session_keys: string[];
    set_header: (key: string, value: string | number | string[]) => void;
    get_header: (key: string) => string | number | string[] | undefined;
    redirect: (code: number, url: string) => void;
    next: () => void;
};
export declare function handle_request(args: ServerHandlerArgs): Promise<void>;
export type Server = {
    use(fn: ServerMiddleware): void;
};
export type ServerMiddleware = (req: IncomingRequest, res: ServerResponse, next: () => void) => void;
export type IncomingRequest = {
    url?: string;
    headers: Headers;
};
export type ServerResponse = {
    redirect(url: string, status?: number): void;
    set_header(name: string, value: string): void;
};
export declare function get_session(req: Headers, secrets: string[]): Promise<App.Session>;
export {};
