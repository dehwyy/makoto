/* eslint-disable */
// @generated by protobuf-ts 2.9.1 with parameter generate_dependencies,eslint_disable
// @generated from protobuf file "hashmap.proto" (package "hashmap", syntax proto3)
// tslint:disable
import type { RpcTransport } from "@protobuf-ts/runtime-rpc";
import type { ServiceInfo } from "@protobuf-ts/runtime-rpc";
import { HashmapRPC } from "./hashmap";
import type { EditItemPayload } from "./hashmap";
import type { IsSuccess } from "./general";
import type { RemoveItemPayload } from "./hashmap";
import type { CreateItemResponse } from "./hashmap";
import type { CreateItemPayload } from "./hashmap";
import type { GetTagsResponse } from "./hashmap";
import type { UserId } from "./general";
import { stackIntercept } from "@protobuf-ts/runtime-rpc";
import type { GetItemsResponse } from "./hashmap";
import type { GetItemsPayload } from "./hashmap";
import type { UnaryCall } from "@protobuf-ts/runtime-rpc";
import type { RpcOptions } from "@protobuf-ts/runtime-rpc";
/**
 * @generated from protobuf service hashmap.HashmapRPC
 */
export interface IHashmapRPCClient {
    /**
     * @generated from protobuf rpc: GetItems(hashmap.GetItemsPayload) returns (hashmap.GetItemsResponse);
     */
    getItems(input: GetItemsPayload, options?: RpcOptions): UnaryCall<GetItemsPayload, GetItemsResponse>;
    /**
     * @generated from protobuf rpc: GetTags(general.UserId) returns (hashmap.GetTagsResponse);
     */
    getTags(input: UserId, options?: RpcOptions): UnaryCall<UserId, GetTagsResponse>;
    /**
     * @generated from protobuf rpc: CreateItem(hashmap.CreateItemPayload) returns (hashmap.CreateItemResponse);
     */
    createItem(input: CreateItemPayload, options?: RpcOptions): UnaryCall<CreateItemPayload, CreateItemResponse>;
    /**
     * @generated from protobuf rpc: RemoveItem(hashmap.RemoveItemPayload) returns (general.IsSuccess);
     */
    removeItem(input: RemoveItemPayload, options?: RpcOptions): UnaryCall<RemoveItemPayload, IsSuccess>;
    /**
     * @generated from protobuf rpc: EditItem(hashmap.EditItemPayload) returns (general.IsSuccess);
     */
    editItem(input: EditItemPayload, options?: RpcOptions): UnaryCall<EditItemPayload, IsSuccess>;
}
/**
 * @generated from protobuf service hashmap.HashmapRPC
 */
export class HashmapRPCClient implements IHashmapRPCClient, ServiceInfo {
    typeName = HashmapRPC.typeName;
    methods = HashmapRPC.methods;
    options = HashmapRPC.options;
    constructor(private readonly _transport: RpcTransport) {
    }
    /**
     * @generated from protobuf rpc: GetItems(hashmap.GetItemsPayload) returns (hashmap.GetItemsResponse);
     */
    getItems(input: GetItemsPayload, options?: RpcOptions): UnaryCall<GetItemsPayload, GetItemsResponse> {
        const method = this.methods[0], opt = this._transport.mergeOptions(options);
        return stackIntercept<GetItemsPayload, GetItemsResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: GetTags(general.UserId) returns (hashmap.GetTagsResponse);
     */
    getTags(input: UserId, options?: RpcOptions): UnaryCall<UserId, GetTagsResponse> {
        const method = this.methods[1], opt = this._transport.mergeOptions(options);
        return stackIntercept<UserId, GetTagsResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: CreateItem(hashmap.CreateItemPayload) returns (hashmap.CreateItemResponse);
     */
    createItem(input: CreateItemPayload, options?: RpcOptions): UnaryCall<CreateItemPayload, CreateItemResponse> {
        const method = this.methods[2], opt = this._transport.mergeOptions(options);
        return stackIntercept<CreateItemPayload, CreateItemResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: RemoveItem(hashmap.RemoveItemPayload) returns (general.IsSuccess);
     */
    removeItem(input: RemoveItemPayload, options?: RpcOptions): UnaryCall<RemoveItemPayload, IsSuccess> {
        const method = this.methods[3], opt = this._transport.mergeOptions(options);
        return stackIntercept<RemoveItemPayload, IsSuccess>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: EditItem(hashmap.EditItemPayload) returns (general.IsSuccess);
     */
    editItem(input: EditItemPayload, options?: RpcOptions): UnaryCall<EditItemPayload, IsSuccess> {
        const method = this.methods[4], opt = this._transport.mergeOptions(options);
        return stackIntercept<EditItemPayload, IsSuccess>("unary", this._transport, method, opt, input);
    }
}
