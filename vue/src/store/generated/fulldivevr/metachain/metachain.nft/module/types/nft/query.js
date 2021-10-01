/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
import { PageRequest, PageResponse, } from "../cosmos/base/query/v1beta1/pagination";
import { Owner, Collection, Denom, BaseNFT } from "../nft/nft";
export const protobufPackage = "metachain.nft";
const baseQuerySupplyRequest = { denomId: "", owner: "" };
export const QuerySupplyRequest = {
    encode(message, writer = Writer.create()) {
        if (message.denomId !== "") {
            writer.uint32(10).string(message.denomId);
        }
        if (message.owner !== "") {
            writer.uint32(18).string(message.owner);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQuerySupplyRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.denomId = reader.string();
                    break;
                case 2:
                    message.owner = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQuerySupplyRequest };
        if (object.denomId !== undefined && object.denomId !== null) {
            message.denomId = String(object.denomId);
        }
        else {
            message.denomId = "";
        }
        if (object.owner !== undefined && object.owner !== null) {
            message.owner = String(object.owner);
        }
        else {
            message.owner = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.denomId !== undefined && (obj.denomId = message.denomId);
        message.owner !== undefined && (obj.owner = message.owner);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQuerySupplyRequest };
        if (object.denomId !== undefined && object.denomId !== null) {
            message.denomId = object.denomId;
        }
        else {
            message.denomId = "";
        }
        if (object.owner !== undefined && object.owner !== null) {
            message.owner = object.owner;
        }
        else {
            message.owner = "";
        }
        return message;
    },
};
const baseQuerySupplyResponse = { amount: 0 };
export const QuerySupplyResponse = {
    encode(message, writer = Writer.create()) {
        if (message.amount !== 0) {
            writer.uint32(8).uint64(message.amount);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQuerySupplyResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.amount = longToNumber(reader.uint64());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQuerySupplyResponse };
        if (object.amount !== undefined && object.amount !== null) {
            message.amount = Number(object.amount);
        }
        else {
            message.amount = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.amount !== undefined && (obj.amount = message.amount);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQuerySupplyResponse };
        if (object.amount !== undefined && object.amount !== null) {
            message.amount = object.amount;
        }
        else {
            message.amount = 0;
        }
        return message;
    },
};
const baseQueryOwnerRequest = { denomId: "", owner: "" };
export const QueryOwnerRequest = {
    encode(message, writer = Writer.create()) {
        if (message.denomId !== "") {
            writer.uint32(10).string(message.denomId);
        }
        if (message.owner !== "") {
            writer.uint32(18).string(message.owner);
        }
        if (message.pagination !== undefined) {
            PageRequest.encode(message.pagination, writer.uint32(26).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryOwnerRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.denomId = reader.string();
                    break;
                case 2:
                    message.owner = reader.string();
                    break;
                case 3:
                    message.pagination = PageRequest.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryOwnerRequest };
        if (object.denomId !== undefined && object.denomId !== null) {
            message.denomId = String(object.denomId);
        }
        else {
            message.denomId = "";
        }
        if (object.owner !== undefined && object.owner !== null) {
            message.owner = String(object.owner);
        }
        else {
            message.owner = "";
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.denomId !== undefined && (obj.denomId = message.denomId);
        message.owner !== undefined && (obj.owner = message.owner);
        message.pagination !== undefined &&
            (obj.pagination = message.pagination
                ? PageRequest.toJSON(message.pagination)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryOwnerRequest };
        if (object.denomId !== undefined && object.denomId !== null) {
            message.denomId = object.denomId;
        }
        else {
            message.denomId = "";
        }
        if (object.owner !== undefined && object.owner !== null) {
            message.owner = object.owner;
        }
        else {
            message.owner = "";
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
};
const baseQueryOwnerResponse = {};
export const QueryOwnerResponse = {
    encode(message, writer = Writer.create()) {
        if (message.owner !== undefined) {
            Owner.encode(message.owner, writer.uint32(10).fork()).ldelim();
        }
        if (message.pagination !== undefined) {
            PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryOwnerResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.owner = Owner.decode(reader, reader.uint32());
                    break;
                case 2:
                    message.pagination = PageResponse.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryOwnerResponse };
        if (object.owner !== undefined && object.owner !== null) {
            message.owner = Owner.fromJSON(object.owner);
        }
        else {
            message.owner = undefined;
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.owner !== undefined &&
            (obj.owner = message.owner ? Owner.toJSON(message.owner) : undefined);
        message.pagination !== undefined &&
            (obj.pagination = message.pagination
                ? PageResponse.toJSON(message.pagination)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryOwnerResponse };
        if (object.owner !== undefined && object.owner !== null) {
            message.owner = Owner.fromPartial(object.owner);
        }
        else {
            message.owner = undefined;
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
};
const baseQueryCollectionRequest = { denomId: "" };
export const QueryCollectionRequest = {
    encode(message, writer = Writer.create()) {
        if (message.denomId !== "") {
            writer.uint32(10).string(message.denomId);
        }
        if (message.pagination !== undefined) {
            PageRequest.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryCollectionRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.denomId = reader.string();
                    break;
                case 2:
                    message.pagination = PageRequest.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryCollectionRequest };
        if (object.denomId !== undefined && object.denomId !== null) {
            message.denomId = String(object.denomId);
        }
        else {
            message.denomId = "";
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.denomId !== undefined && (obj.denomId = message.denomId);
        message.pagination !== undefined &&
            (obj.pagination = message.pagination
                ? PageRequest.toJSON(message.pagination)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryCollectionRequest };
        if (object.denomId !== undefined && object.denomId !== null) {
            message.denomId = object.denomId;
        }
        else {
            message.denomId = "";
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
};
const baseQueryCollectionResponse = {};
export const QueryCollectionResponse = {
    encode(message, writer = Writer.create()) {
        if (message.collection !== undefined) {
            Collection.encode(message.collection, writer.uint32(10).fork()).ldelim();
        }
        if (message.pagination !== undefined) {
            PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQueryCollectionResponse,
        };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.collection = Collection.decode(reader, reader.uint32());
                    break;
                case 2:
                    message.pagination = PageResponse.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = {
            ...baseQueryCollectionResponse,
        };
        if (object.collection !== undefined && object.collection !== null) {
            message.collection = Collection.fromJSON(object.collection);
        }
        else {
            message.collection = undefined;
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.collection !== undefined &&
            (obj.collection = message.collection
                ? Collection.toJSON(message.collection)
                : undefined);
        message.pagination !== undefined &&
            (obj.pagination = message.pagination
                ? PageResponse.toJSON(message.pagination)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryCollectionResponse,
        };
        if (object.collection !== undefined && object.collection !== null) {
            message.collection = Collection.fromPartial(object.collection);
        }
        else {
            message.collection = undefined;
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
};
const baseQueryDenomRequest = { denomId: "" };
export const QueryDenomRequest = {
    encode(message, writer = Writer.create()) {
        if (message.denomId !== "") {
            writer.uint32(10).string(message.denomId);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryDenomRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.denomId = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryDenomRequest };
        if (object.denomId !== undefined && object.denomId !== null) {
            message.denomId = String(object.denomId);
        }
        else {
            message.denomId = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.denomId !== undefined && (obj.denomId = message.denomId);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryDenomRequest };
        if (object.denomId !== undefined && object.denomId !== null) {
            message.denomId = object.denomId;
        }
        else {
            message.denomId = "";
        }
        return message;
    },
};
const baseQueryDenomResponse = {};
export const QueryDenomResponse = {
    encode(message, writer = Writer.create()) {
        if (message.denom !== undefined) {
            Denom.encode(message.denom, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryDenomResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.denom = Denom.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryDenomResponse };
        if (object.denom !== undefined && object.denom !== null) {
            message.denom = Denom.fromJSON(object.denom);
        }
        else {
            message.denom = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.denom !== undefined &&
            (obj.denom = message.denom ? Denom.toJSON(message.denom) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryDenomResponse };
        if (object.denom !== undefined && object.denom !== null) {
            message.denom = Denom.fromPartial(object.denom);
        }
        else {
            message.denom = undefined;
        }
        return message;
    },
};
const baseQueryDenomsRequest = {};
export const QueryDenomsRequest = {
    encode(message, writer = Writer.create()) {
        if (message.pagination !== undefined) {
            PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryDenomsRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.pagination = PageRequest.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryDenomsRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.pagination !== undefined &&
            (obj.pagination = message.pagination
                ? PageRequest.toJSON(message.pagination)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryDenomsRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
};
const baseQueryDenomsResponse = {};
export const QueryDenomsResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.denoms) {
            Denom.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.pagination !== undefined) {
            PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryDenomsResponse };
        message.denoms = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.denoms.push(Denom.decode(reader, reader.uint32()));
                    break;
                case 2:
                    message.pagination = PageResponse.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryDenomsResponse };
        message.denoms = [];
        if (object.denoms !== undefined && object.denoms !== null) {
            for (const e of object.denoms) {
                message.denoms.push(Denom.fromJSON(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.denoms) {
            obj.denoms = message.denoms.map((e) => (e ? Denom.toJSON(e) : undefined));
        }
        else {
            obj.denoms = [];
        }
        message.pagination !== undefined &&
            (obj.pagination = message.pagination
                ? PageResponse.toJSON(message.pagination)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryDenomsResponse };
        message.denoms = [];
        if (object.denoms !== undefined && object.denoms !== null) {
            for (const e of object.denoms) {
                message.denoms.push(Denom.fromPartial(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
};
const baseQueryNFTRequest = { denomId: "", tokenId: "" };
export const QueryNFTRequest = {
    encode(message, writer = Writer.create()) {
        if (message.denomId !== "") {
            writer.uint32(10).string(message.denomId);
        }
        if (message.tokenId !== "") {
            writer.uint32(18).string(message.tokenId);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryNFTRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.denomId = reader.string();
                    break;
                case 2:
                    message.tokenId = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryNFTRequest };
        if (object.denomId !== undefined && object.denomId !== null) {
            message.denomId = String(object.denomId);
        }
        else {
            message.denomId = "";
        }
        if (object.tokenId !== undefined && object.tokenId !== null) {
            message.tokenId = String(object.tokenId);
        }
        else {
            message.tokenId = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.denomId !== undefined && (obj.denomId = message.denomId);
        message.tokenId !== undefined && (obj.tokenId = message.tokenId);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryNFTRequest };
        if (object.denomId !== undefined && object.denomId !== null) {
            message.denomId = object.denomId;
        }
        else {
            message.denomId = "";
        }
        if (object.tokenId !== undefined && object.tokenId !== null) {
            message.tokenId = object.tokenId;
        }
        else {
            message.tokenId = "";
        }
        return message;
    },
};
const baseQueryNFTResponse = {};
export const QueryNFTResponse = {
    encode(message, writer = Writer.create()) {
        if (message.nft !== undefined) {
            BaseNFT.encode(message.nft, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryNFTResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.nft = BaseNFT.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryNFTResponse };
        if (object.nft !== undefined && object.nft !== null) {
            message.nft = BaseNFT.fromJSON(object.nft);
        }
        else {
            message.nft = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.nft !== undefined &&
            (obj.nft = message.nft ? BaseNFT.toJSON(message.nft) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryNFTResponse };
        if (object.nft !== undefined && object.nft !== null) {
            message.nft = BaseNFT.fromPartial(object.nft);
        }
        else {
            message.nft = undefined;
        }
        return message;
    },
};
export class QueryClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
    }
    Supply(request) {
        const data = QuerySupplyRequest.encode(request).finish();
        const promise = this.rpc.request("metachain.nft.Query", "Supply", data);
        return promise.then((data) => QuerySupplyResponse.decode(new Reader(data)));
    }
    Owner(request) {
        const data = QueryOwnerRequest.encode(request).finish();
        const promise = this.rpc.request("metachain.nft.Query", "Owner", data);
        return promise.then((data) => QueryOwnerResponse.decode(new Reader(data)));
    }
    Collection(request) {
        const data = QueryCollectionRequest.encode(request).finish();
        const promise = this.rpc.request("metachain.nft.Query", "Collection", data);
        return promise.then((data) => QueryCollectionResponse.decode(new Reader(data)));
    }
    Denom(request) {
        const data = QueryDenomRequest.encode(request).finish();
        const promise = this.rpc.request("metachain.nft.Query", "Denom", data);
        return promise.then((data) => QueryDenomResponse.decode(new Reader(data)));
    }
    Denoms(request) {
        const data = QueryDenomsRequest.encode(request).finish();
        const promise = this.rpc.request("metachain.nft.Query", "Denoms", data);
        return promise.then((data) => QueryDenomsResponse.decode(new Reader(data)));
    }
    Nft(request) {
        const data = QueryNFTRequest.encode(request).finish();
        const promise = this.rpc.request("metachain.nft.Query", "Nft", data);
        return promise.then((data) => QueryNFTResponse.decode(new Reader(data)));
    }
}
var globalThis = (() => {
    if (typeof globalThis !== "undefined")
        return globalThis;
    if (typeof self !== "undefined")
        return self;
    if (typeof window !== "undefined")
        return window;
    if (typeof global !== "undefined")
        return global;
    throw "Unable to locate global object";
})();
function longToNumber(long) {
    if (long.gt(Number.MAX_SAFE_INTEGER)) {
        throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
    }
    return long.toNumber();
}
if (util.Long !== Long) {
    util.Long = Long;
    configure();
}
