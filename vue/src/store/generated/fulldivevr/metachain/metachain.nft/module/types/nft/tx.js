/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
export const protobufPackage = "metachain.nft";
const baseMsgIssueDenom = {
    id: "",
    name: "",
    schema: "",
    sender: "",
    symbol: "",
    mintRestricted: false,
    updateRestricted: false,
};
export const MsgIssueDenom = {
    encode(message, writer = Writer.create()) {
        if (message.id !== "") {
            writer.uint32(10).string(message.id);
        }
        if (message.name !== "") {
            writer.uint32(18).string(message.name);
        }
        if (message.schema !== "") {
            writer.uint32(26).string(message.schema);
        }
        if (message.sender !== "") {
            writer.uint32(34).string(message.sender);
        }
        if (message.symbol !== "") {
            writer.uint32(42).string(message.symbol);
        }
        if (message.mintRestricted === true) {
            writer.uint32(48).bool(message.mintRestricted);
        }
        if (message.updateRestricted === true) {
            writer.uint32(56).bool(message.updateRestricted);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgIssueDenom };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.id = reader.string();
                    break;
                case 2:
                    message.name = reader.string();
                    break;
                case 3:
                    message.schema = reader.string();
                    break;
                case 4:
                    message.sender = reader.string();
                    break;
                case 5:
                    message.symbol = reader.string();
                    break;
                case 6:
                    message.mintRestricted = reader.bool();
                    break;
                case 7:
                    message.updateRestricted = reader.bool();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgIssueDenom };
        if (object.id !== undefined && object.id !== null) {
            message.id = String(object.id);
        }
        else {
            message.id = "";
        }
        if (object.name !== undefined && object.name !== null) {
            message.name = String(object.name);
        }
        else {
            message.name = "";
        }
        if (object.schema !== undefined && object.schema !== null) {
            message.schema = String(object.schema);
        }
        else {
            message.schema = "";
        }
        if (object.sender !== undefined && object.sender !== null) {
            message.sender = String(object.sender);
        }
        else {
            message.sender = "";
        }
        if (object.symbol !== undefined && object.symbol !== null) {
            message.symbol = String(object.symbol);
        }
        else {
            message.symbol = "";
        }
        if (object.mintRestricted !== undefined && object.mintRestricted !== null) {
            message.mintRestricted = Boolean(object.mintRestricted);
        }
        else {
            message.mintRestricted = false;
        }
        if (object.updateRestricted !== undefined &&
            object.updateRestricted !== null) {
            message.updateRestricted = Boolean(object.updateRestricted);
        }
        else {
            message.updateRestricted = false;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.id !== undefined && (obj.id = message.id);
        message.name !== undefined && (obj.name = message.name);
        message.schema !== undefined && (obj.schema = message.schema);
        message.sender !== undefined && (obj.sender = message.sender);
        message.symbol !== undefined && (obj.symbol = message.symbol);
        message.mintRestricted !== undefined &&
            (obj.mintRestricted = message.mintRestricted);
        message.updateRestricted !== undefined &&
            (obj.updateRestricted = message.updateRestricted);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgIssueDenom };
        if (object.id !== undefined && object.id !== null) {
            message.id = object.id;
        }
        else {
            message.id = "";
        }
        if (object.name !== undefined && object.name !== null) {
            message.name = object.name;
        }
        else {
            message.name = "";
        }
        if (object.schema !== undefined && object.schema !== null) {
            message.schema = object.schema;
        }
        else {
            message.schema = "";
        }
        if (object.sender !== undefined && object.sender !== null) {
            message.sender = object.sender;
        }
        else {
            message.sender = "";
        }
        if (object.symbol !== undefined && object.symbol !== null) {
            message.symbol = object.symbol;
        }
        else {
            message.symbol = "";
        }
        if (object.mintRestricted !== undefined && object.mintRestricted !== null) {
            message.mintRestricted = object.mintRestricted;
        }
        else {
            message.mintRestricted = false;
        }
        if (object.updateRestricted !== undefined &&
            object.updateRestricted !== null) {
            message.updateRestricted = object.updateRestricted;
        }
        else {
            message.updateRestricted = false;
        }
        return message;
    },
};
const baseMsgIssueDenomResponse = {};
export const MsgIssueDenomResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgIssueDenomResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(_) {
        const message = { ...baseMsgIssueDenomResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgIssueDenomResponse };
        return message;
    },
};
const baseMsgTransferNFT = {
    id: "",
    denomId: "",
    name: "",
    uri: "",
    data: "",
    sender: "",
    recipient: "",
};
export const MsgTransferNFT = {
    encode(message, writer = Writer.create()) {
        if (message.id !== "") {
            writer.uint32(10).string(message.id);
        }
        if (message.denomId !== "") {
            writer.uint32(18).string(message.denomId);
        }
        if (message.name !== "") {
            writer.uint32(26).string(message.name);
        }
        if (message.uri !== "") {
            writer.uint32(34).string(message.uri);
        }
        if (message.data !== "") {
            writer.uint32(42).string(message.data);
        }
        if (message.sender !== "") {
            writer.uint32(50).string(message.sender);
        }
        if (message.recipient !== "") {
            writer.uint32(58).string(message.recipient);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgTransferNFT };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.id = reader.string();
                    break;
                case 2:
                    message.denomId = reader.string();
                    break;
                case 3:
                    message.name = reader.string();
                    break;
                case 4:
                    message.uri = reader.string();
                    break;
                case 5:
                    message.data = reader.string();
                    break;
                case 6:
                    message.sender = reader.string();
                    break;
                case 7:
                    message.recipient = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgTransferNFT };
        if (object.id !== undefined && object.id !== null) {
            message.id = String(object.id);
        }
        else {
            message.id = "";
        }
        if (object.denomId !== undefined && object.denomId !== null) {
            message.denomId = String(object.denomId);
        }
        else {
            message.denomId = "";
        }
        if (object.name !== undefined && object.name !== null) {
            message.name = String(object.name);
        }
        else {
            message.name = "";
        }
        if (object.uri !== undefined && object.uri !== null) {
            message.uri = String(object.uri);
        }
        else {
            message.uri = "";
        }
        if (object.data !== undefined && object.data !== null) {
            message.data = String(object.data);
        }
        else {
            message.data = "";
        }
        if (object.sender !== undefined && object.sender !== null) {
            message.sender = String(object.sender);
        }
        else {
            message.sender = "";
        }
        if (object.recipient !== undefined && object.recipient !== null) {
            message.recipient = String(object.recipient);
        }
        else {
            message.recipient = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.id !== undefined && (obj.id = message.id);
        message.denomId !== undefined && (obj.denomId = message.denomId);
        message.name !== undefined && (obj.name = message.name);
        message.uri !== undefined && (obj.uri = message.uri);
        message.data !== undefined && (obj.data = message.data);
        message.sender !== undefined && (obj.sender = message.sender);
        message.recipient !== undefined && (obj.recipient = message.recipient);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgTransferNFT };
        if (object.id !== undefined && object.id !== null) {
            message.id = object.id;
        }
        else {
            message.id = "";
        }
        if (object.denomId !== undefined && object.denomId !== null) {
            message.denomId = object.denomId;
        }
        else {
            message.denomId = "";
        }
        if (object.name !== undefined && object.name !== null) {
            message.name = object.name;
        }
        else {
            message.name = "";
        }
        if (object.uri !== undefined && object.uri !== null) {
            message.uri = object.uri;
        }
        else {
            message.uri = "";
        }
        if (object.data !== undefined && object.data !== null) {
            message.data = object.data;
        }
        else {
            message.data = "";
        }
        if (object.sender !== undefined && object.sender !== null) {
            message.sender = object.sender;
        }
        else {
            message.sender = "";
        }
        if (object.recipient !== undefined && object.recipient !== null) {
            message.recipient = object.recipient;
        }
        else {
            message.recipient = "";
        }
        return message;
    },
};
const baseMsgTransferNFTResponse = {};
export const MsgTransferNFTResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgTransferNFTResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(_) {
        const message = { ...baseMsgTransferNFTResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgTransferNFTResponse };
        return message;
    },
};
const baseMsgEditNFT = {
    id: "",
    denomId: "",
    name: "",
    uri: "",
    data: "",
    sender: "",
};
export const MsgEditNFT = {
    encode(message, writer = Writer.create()) {
        if (message.id !== "") {
            writer.uint32(10).string(message.id);
        }
        if (message.denomId !== "") {
            writer.uint32(18).string(message.denomId);
        }
        if (message.name !== "") {
            writer.uint32(26).string(message.name);
        }
        if (message.uri !== "") {
            writer.uint32(34).string(message.uri);
        }
        if (message.data !== "") {
            writer.uint32(42).string(message.data);
        }
        if (message.sender !== "") {
            writer.uint32(50).string(message.sender);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgEditNFT };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.id = reader.string();
                    break;
                case 2:
                    message.denomId = reader.string();
                    break;
                case 3:
                    message.name = reader.string();
                    break;
                case 4:
                    message.uri = reader.string();
                    break;
                case 5:
                    message.data = reader.string();
                    break;
                case 6:
                    message.sender = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgEditNFT };
        if (object.id !== undefined && object.id !== null) {
            message.id = String(object.id);
        }
        else {
            message.id = "";
        }
        if (object.denomId !== undefined && object.denomId !== null) {
            message.denomId = String(object.denomId);
        }
        else {
            message.denomId = "";
        }
        if (object.name !== undefined && object.name !== null) {
            message.name = String(object.name);
        }
        else {
            message.name = "";
        }
        if (object.uri !== undefined && object.uri !== null) {
            message.uri = String(object.uri);
        }
        else {
            message.uri = "";
        }
        if (object.data !== undefined && object.data !== null) {
            message.data = String(object.data);
        }
        else {
            message.data = "";
        }
        if (object.sender !== undefined && object.sender !== null) {
            message.sender = String(object.sender);
        }
        else {
            message.sender = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.id !== undefined && (obj.id = message.id);
        message.denomId !== undefined && (obj.denomId = message.denomId);
        message.name !== undefined && (obj.name = message.name);
        message.uri !== undefined && (obj.uri = message.uri);
        message.data !== undefined && (obj.data = message.data);
        message.sender !== undefined && (obj.sender = message.sender);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgEditNFT };
        if (object.id !== undefined && object.id !== null) {
            message.id = object.id;
        }
        else {
            message.id = "";
        }
        if (object.denomId !== undefined && object.denomId !== null) {
            message.denomId = object.denomId;
        }
        else {
            message.denomId = "";
        }
        if (object.name !== undefined && object.name !== null) {
            message.name = object.name;
        }
        else {
            message.name = "";
        }
        if (object.uri !== undefined && object.uri !== null) {
            message.uri = object.uri;
        }
        else {
            message.uri = "";
        }
        if (object.data !== undefined && object.data !== null) {
            message.data = object.data;
        }
        else {
            message.data = "";
        }
        if (object.sender !== undefined && object.sender !== null) {
            message.sender = object.sender;
        }
        else {
            message.sender = "";
        }
        return message;
    },
};
const baseMsgEditNFTResponse = {};
export const MsgEditNFTResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgEditNFTResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(_) {
        const message = { ...baseMsgEditNFTResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgEditNFTResponse };
        return message;
    },
};
const baseMsgMintNFT = {
    id: "",
    denomId: "",
    name: "",
    uri: "",
    data: "",
    sender: "",
    recipient: "",
};
export const MsgMintNFT = {
    encode(message, writer = Writer.create()) {
        if (message.id !== "") {
            writer.uint32(10).string(message.id);
        }
        if (message.denomId !== "") {
            writer.uint32(18).string(message.denomId);
        }
        if (message.name !== "") {
            writer.uint32(26).string(message.name);
        }
        if (message.uri !== "") {
            writer.uint32(34).string(message.uri);
        }
        if (message.data !== "") {
            writer.uint32(42).string(message.data);
        }
        if (message.sender !== "") {
            writer.uint32(50).string(message.sender);
        }
        if (message.recipient !== "") {
            writer.uint32(58).string(message.recipient);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgMintNFT };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.id = reader.string();
                    break;
                case 2:
                    message.denomId = reader.string();
                    break;
                case 3:
                    message.name = reader.string();
                    break;
                case 4:
                    message.uri = reader.string();
                    break;
                case 5:
                    message.data = reader.string();
                    break;
                case 6:
                    message.sender = reader.string();
                    break;
                case 7:
                    message.recipient = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgMintNFT };
        if (object.id !== undefined && object.id !== null) {
            message.id = String(object.id);
        }
        else {
            message.id = "";
        }
        if (object.denomId !== undefined && object.denomId !== null) {
            message.denomId = String(object.denomId);
        }
        else {
            message.denomId = "";
        }
        if (object.name !== undefined && object.name !== null) {
            message.name = String(object.name);
        }
        else {
            message.name = "";
        }
        if (object.uri !== undefined && object.uri !== null) {
            message.uri = String(object.uri);
        }
        else {
            message.uri = "";
        }
        if (object.data !== undefined && object.data !== null) {
            message.data = String(object.data);
        }
        else {
            message.data = "";
        }
        if (object.sender !== undefined && object.sender !== null) {
            message.sender = String(object.sender);
        }
        else {
            message.sender = "";
        }
        if (object.recipient !== undefined && object.recipient !== null) {
            message.recipient = String(object.recipient);
        }
        else {
            message.recipient = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.id !== undefined && (obj.id = message.id);
        message.denomId !== undefined && (obj.denomId = message.denomId);
        message.name !== undefined && (obj.name = message.name);
        message.uri !== undefined && (obj.uri = message.uri);
        message.data !== undefined && (obj.data = message.data);
        message.sender !== undefined && (obj.sender = message.sender);
        message.recipient !== undefined && (obj.recipient = message.recipient);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgMintNFT };
        if (object.id !== undefined && object.id !== null) {
            message.id = object.id;
        }
        else {
            message.id = "";
        }
        if (object.denomId !== undefined && object.denomId !== null) {
            message.denomId = object.denomId;
        }
        else {
            message.denomId = "";
        }
        if (object.name !== undefined && object.name !== null) {
            message.name = object.name;
        }
        else {
            message.name = "";
        }
        if (object.uri !== undefined && object.uri !== null) {
            message.uri = object.uri;
        }
        else {
            message.uri = "";
        }
        if (object.data !== undefined && object.data !== null) {
            message.data = object.data;
        }
        else {
            message.data = "";
        }
        if (object.sender !== undefined && object.sender !== null) {
            message.sender = object.sender;
        }
        else {
            message.sender = "";
        }
        if (object.recipient !== undefined && object.recipient !== null) {
            message.recipient = object.recipient;
        }
        else {
            message.recipient = "";
        }
        return message;
    },
};
const baseMsgMintNFTResponse = {};
export const MsgMintNFTResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgMintNFTResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(_) {
        const message = { ...baseMsgMintNFTResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgMintNFTResponse };
        return message;
    },
};
const baseMsgBurnNFT = { id: "", denomId: "", sender: "" };
export const MsgBurnNFT = {
    encode(message, writer = Writer.create()) {
        if (message.id !== "") {
            writer.uint32(10).string(message.id);
        }
        if (message.denomId !== "") {
            writer.uint32(18).string(message.denomId);
        }
        if (message.sender !== "") {
            writer.uint32(26).string(message.sender);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgBurnNFT };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.id = reader.string();
                    break;
                case 2:
                    message.denomId = reader.string();
                    break;
                case 3:
                    message.sender = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgBurnNFT };
        if (object.id !== undefined && object.id !== null) {
            message.id = String(object.id);
        }
        else {
            message.id = "";
        }
        if (object.denomId !== undefined && object.denomId !== null) {
            message.denomId = String(object.denomId);
        }
        else {
            message.denomId = "";
        }
        if (object.sender !== undefined && object.sender !== null) {
            message.sender = String(object.sender);
        }
        else {
            message.sender = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.id !== undefined && (obj.id = message.id);
        message.denomId !== undefined && (obj.denomId = message.denomId);
        message.sender !== undefined && (obj.sender = message.sender);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgBurnNFT };
        if (object.id !== undefined && object.id !== null) {
            message.id = object.id;
        }
        else {
            message.id = "";
        }
        if (object.denomId !== undefined && object.denomId !== null) {
            message.denomId = object.denomId;
        }
        else {
            message.denomId = "";
        }
        if (object.sender !== undefined && object.sender !== null) {
            message.sender = object.sender;
        }
        else {
            message.sender = "";
        }
        return message;
    },
};
const baseMsgBurnNFTResponse = {};
export const MsgBurnNFTResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgBurnNFTResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(_) {
        const message = { ...baseMsgBurnNFTResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgBurnNFTResponse };
        return message;
    },
};
const baseMsgTransferDenom = { id: "", sender: "", recipient: "" };
export const MsgTransferDenom = {
    encode(message, writer = Writer.create()) {
        if (message.id !== "") {
            writer.uint32(10).string(message.id);
        }
        if (message.sender !== "") {
            writer.uint32(18).string(message.sender);
        }
        if (message.recipient !== "") {
            writer.uint32(26).string(message.recipient);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgTransferDenom };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.id = reader.string();
                    break;
                case 2:
                    message.sender = reader.string();
                    break;
                case 3:
                    message.recipient = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgTransferDenom };
        if (object.id !== undefined && object.id !== null) {
            message.id = String(object.id);
        }
        else {
            message.id = "";
        }
        if (object.sender !== undefined && object.sender !== null) {
            message.sender = String(object.sender);
        }
        else {
            message.sender = "";
        }
        if (object.recipient !== undefined && object.recipient !== null) {
            message.recipient = String(object.recipient);
        }
        else {
            message.recipient = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.id !== undefined && (obj.id = message.id);
        message.sender !== undefined && (obj.sender = message.sender);
        message.recipient !== undefined && (obj.recipient = message.recipient);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgTransferDenom };
        if (object.id !== undefined && object.id !== null) {
            message.id = object.id;
        }
        else {
            message.id = "";
        }
        if (object.sender !== undefined && object.sender !== null) {
            message.sender = object.sender;
        }
        else {
            message.sender = "";
        }
        if (object.recipient !== undefined && object.recipient !== null) {
            message.recipient = object.recipient;
        }
        else {
            message.recipient = "";
        }
        return message;
    },
};
const baseMsgTransferDenomResponse = {};
export const MsgTransferDenomResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseMsgTransferDenomResponse,
        };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(_) {
        const message = {
            ...baseMsgTransferDenomResponse,
        };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = {
            ...baseMsgTransferDenomResponse,
        };
        return message;
    },
};
export class MsgClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
    }
    IssueDenom(request) {
        const data = MsgIssueDenom.encode(request).finish();
        const promise = this.rpc.request("metachain.nft.Msg", "IssueDenom", data);
        return promise.then((data) => MsgIssueDenomResponse.decode(new Reader(data)));
    }
    MintNFT(request) {
        const data = MsgMintNFT.encode(request).finish();
        const promise = this.rpc.request("metachain.nft.Msg", "MintNFT", data);
        return promise.then((data) => MsgMintNFTResponse.decode(new Reader(data)));
    }
    EditNFT(request) {
        const data = MsgEditNFT.encode(request).finish();
        const promise = this.rpc.request("metachain.nft.Msg", "EditNFT", data);
        return promise.then((data) => MsgEditNFTResponse.decode(new Reader(data)));
    }
    TransferNFT(request) {
        const data = MsgTransferNFT.encode(request).finish();
        const promise = this.rpc.request("metachain.nft.Msg", "TransferNFT", data);
        return promise.then((data) => MsgTransferNFTResponse.decode(new Reader(data)));
    }
    BurnNFT(request) {
        const data = MsgBurnNFT.encode(request).finish();
        const promise = this.rpc.request("metachain.nft.Msg", "BurnNFT", data);
        return promise.then((data) => MsgBurnNFTResponse.decode(new Reader(data)));
    }
    TransferDenom(request) {
        const data = MsgTransferDenom.encode(request).finish();
        const promise = this.rpc.request("metachain.nft.Msg", "TransferDenom", data);
        return promise.then((data) => MsgTransferDenomResponse.decode(new Reader(data)));
    }
}
