export class RpcInterceptors {
    static WithToken(cookies, key = "token") {
        return this.AddAuthorizationHeader(cookies.get(key));
    }
    static AddAuthorizationHeader(header_value) {
        return {
            interceptUnary(next, method, input, options) {
                if (!options.meta) {
                    options.meta = {};
                }
                if (header_value) {
                    options.meta['Authorization'] = 'Bearer ' + header_value;
                }
                return next(method, input, options);
            }
        };
    }
}
