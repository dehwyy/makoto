import type { Record } from "./public/record";

type GetUserByIdInput = {
    userId: string;
};

type GetUserByUsernameInput = {
    username: string;
};

export declare type CacheTypeDef = {
    types: {
        __ROOT__: {
            idFields: {};
            fields: {
                getQuestion: {
                    type: Record<CacheTypeDef, "UserQuestionResponse">;
                    args: never;
                };
                getUserById: {
                    type: Record<CacheTypeDef, "UserResponse">;
                    args: {
                        input: GetUserByIdInput;
                    };
                };
                getUserByUsername: {
                    type: Record<CacheTypeDef, "UserResponse">;
                    args: {
                        input: GetUserByUsernameInput;
                    };
                };
            };
            fragments: [];
        };
        Status: {
            idFields: never;
            fields: {
                is_ok: {
                    type: boolean;
                    args: never;
                };
            };
            fragments: [];
        };
        Tokens: {
            idFields: never;
            fields: {
                access_token: {
                    type: string;
                    args: never;
                };
                refresh_token: {
                    type: string;
                    args: never;
                };
            };
            fragments: [];
        };
        UserAuthResponse: {
            idFields: never;
            fields: {
                tokens: {
                    type: Record<CacheTypeDef, "Tokens">;
                    args: never;
                };
                userId: {
                    type: string;
                    args: never;
                };
            };
            fragments: [];
        };
        UserQuestionResponse: {
            idFields: never;
            fields: {
                auth: {
                    type: Record<CacheTypeDef, "UserAuthResponse">;
                    args: never;
                };
                question: {
                    type: string;
                    args: never;
                };
            };
            fragments: [];
        };
        UserResponse: {
            idFields: never;
            fields: {
                auth: {
                    type: Record<CacheTypeDef, "UserAuthResponse">;
                    args: never;
                };
                username: {
                    type: string;
                    args: never;
                };
            };
            fragments: [];
        };
    };
    lists: {};
    queries: [];
};