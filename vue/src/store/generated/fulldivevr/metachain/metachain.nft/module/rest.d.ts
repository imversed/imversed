export interface NftBaseNFT {
    id?: string;
    name?: string;
    uri?: string;
    data?: string;
    owner?: string;
}
export interface NftCollection {
    denom?: NftDenom;
    nfts?: NftBaseNFT[];
}
export interface NftDenom {
    id?: string;
    name?: string;
    schema?: string;
    creator?: string;
    symbol?: string;
    mintRestricted?: boolean;
    updateRestricted?: boolean;
}
export interface NftIDCollection {
    denomId?: string;
    tokenIds?: string[];
}
/**
 * MsgBurnNFTResponse defines the Msg/BurnNFT response type.
 */
export declare type NftMsgBurnNFTResponse = object;
/**
 * MsgEditNFTResponse defines the Msg/EditNFT response type.
 */
export declare type NftMsgEditNFTResponse = object;
/**
 * MsgIssueDenomResponse defines the Msg/IssueDenom response type.
 */
export declare type NftMsgIssueDenomResponse = object;
/**
 * MsgMintNFTResponse defines the Msg/MintNFT response type.
 */
export declare type NftMsgMintNFTResponse = object;
/**
 * MsgTransferDenomResponse defines the Msg/TransferDenom response type.
 */
export declare type NftMsgTransferDenomResponse = object;
/**
 * MsgTransferNFTResponse defines the Msg/TransferNFT response type.
 */
export declare type NftMsgTransferNFTResponse = object;
export interface NftOwner {
    address?: string;
    idCollections?: NftIDCollection[];
}
export interface NftQueryCollectionResponse {
    collection?: NftCollection;
    /**
     * PageResponse is to be embedded in gRPC response messages where the
     * corresponding request message has used PageRequest.
     *
     *  message SomeResponse {
     *          repeated Bar results = 1;
     *          PageResponse page = 2;
     *  }
     */
    pagination?: V1Beta1PageResponse;
}
export interface NftQueryDenomResponse {
    denom?: NftDenom;
}
export interface NftQueryDenomsResponse {
    denoms?: NftDenom[];
    /**
     * PageResponse is to be embedded in gRPC response messages where the
     * corresponding request message has used PageRequest.
     *
     *  message SomeResponse {
     *          repeated Bar results = 1;
     *          PageResponse page = 2;
     *  }
     */
    pagination?: V1Beta1PageResponse;
}
export interface NftQueryNFTResponse {
    nft?: NftBaseNFT;
}
export interface NftQueryOwnerResponse {
    owner?: NftOwner;
    /**
     * PageResponse is to be embedded in gRPC response messages where the
     * corresponding request message has used PageRequest.
     *
     *  message SomeResponse {
     *          repeated Bar results = 1;
     *          PageResponse page = 2;
     *  }
     */
    pagination?: V1Beta1PageResponse;
}
export interface NftQuerySupplyResponse {
    /** @format uint64 */
    amount?: string;
}
export interface ProtobufAny {
    "@type"?: string;
}
export interface RpcStatus {
    /** @format int32 */
    code?: number;
    message?: string;
    details?: ProtobufAny[];
}
/**
* message SomeRequest {
         Foo some_parameter = 1;
         PageRequest pagination = 2;
 }
*/
export interface V1Beta1PageRequest {
    /**
     * key is a value returned in PageResponse.next_key to begin
     * querying the next page most efficiently. Only one of offset or key
     * should be set.
     * @format byte
     */
    key?: string;
    /**
     * offset is a numeric offset that can be used when key is unavailable.
     * It is less efficient than using key. Only one of offset or key should
     * be set.
     * @format uint64
     */
    offset?: string;
    /**
     * limit is the total number of results to be returned in the result page.
     * If left empty it will default to a value to be set by each app.
     * @format uint64
     */
    limit?: string;
    /**
     * count_total is set to true  to indicate that the result set should include
     * a count of the total number of items available for pagination in UIs.
     * count_total is only respected when offset is used. It is ignored when key
     * is set.
     */
    countTotal?: boolean;
    /** reverse is set to true if results are to be returned in the descending order. */
    reverse?: boolean;
}
/**
* PageResponse is to be embedded in gRPC response messages where the
corresponding request message has used PageRequest.

 message SomeResponse {
         repeated Bar results = 1;
         PageResponse page = 2;
 }
*/
export interface V1Beta1PageResponse {
    /** @format byte */
    nextKey?: string;
    /** @format uint64 */
    total?: string;
}
export declare type QueryParamsType = Record<string | number, any>;
export declare type ResponseFormat = keyof Omit<Body, "body" | "bodyUsed">;
export interface FullRequestParams extends Omit<RequestInit, "body"> {
    /** set parameter to `true` for call `securityWorker` for this request */
    secure?: boolean;
    /** request path */
    path: string;
    /** content type of request body */
    type?: ContentType;
    /** query params */
    query?: QueryParamsType;
    /** format of response (i.e. response.json() -> format: "json") */
    format?: keyof Omit<Body, "body" | "bodyUsed">;
    /** request body */
    body?: unknown;
    /** base url */
    baseUrl?: string;
    /** request cancellation token */
    cancelToken?: CancelToken;
}
export declare type RequestParams = Omit<FullRequestParams, "body" | "method" | "query" | "path">;
export interface ApiConfig<SecurityDataType = unknown> {
    baseUrl?: string;
    baseApiParams?: Omit<RequestParams, "baseUrl" | "cancelToken" | "signal">;
    securityWorker?: (securityData: SecurityDataType) => RequestParams | void;
}
export interface HttpResponse<D extends unknown, E extends unknown = unknown> extends Response {
    data: D;
    error: E;
}
declare type CancelToken = Symbol | string | number;
export declare enum ContentType {
    Json = "application/json",
    FormData = "multipart/form-data",
    UrlEncoded = "application/x-www-form-urlencoded"
}
export declare class HttpClient<SecurityDataType = unknown> {
    baseUrl: string;
    private securityData;
    private securityWorker;
    private abortControllers;
    private baseApiParams;
    constructor(apiConfig?: ApiConfig<SecurityDataType>);
    setSecurityData: (data: SecurityDataType) => void;
    private addQueryParam;
    protected toQueryString(rawQuery?: QueryParamsType): string;
    protected addQueryParams(rawQuery?: QueryParamsType): string;
    private contentFormatters;
    private mergeRequestParams;
    private createAbortSignal;
    abortRequest: (cancelToken: CancelToken) => void;
    request: <T = any, E = any>({ body, secure, path, type, query, format, baseUrl, cancelToken, ...params }: FullRequestParams) => Promise<HttpResponse<T, E>>;
}
/**
 * @title nft/genesis.proto
 * @version version not set
 */
export declare class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
    /**
     * No description
     *
     * @tags Query
     * @name QueryCollection
     * @summary Collection queries the NFTs of the specified denom
     * @request GET:/metachain/nft/collections/{denomId}
     */
    queryCollection: (denomId: string, query?: {
        "pagination.key"?: string;
        "pagination.offset"?: string;
        "pagination.limit"?: string;
        "pagination.countTotal"?: boolean;
        "pagination.reverse"?: boolean;
    }, params?: RequestParams) => Promise<HttpResponse<NftQueryCollectionResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QuerySupply
     * @summary Supply queries the total supply of a given denom or owner
     * @request GET:/metachain/nft/collections/{denomId}/supply
     */
    querySupply: (denomId: string, query?: {
        owner?: string;
    }, params?: RequestParams) => Promise<HttpResponse<NftQuerySupplyResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryDenoms
     * @summary Denoms queries all the denoms
     * @request GET:/metachain/nft/denoms
     */
    queryDenoms: (query?: {
        "pagination.key"?: string;
        "pagination.offset"?: string;
        "pagination.limit"?: string;
        "pagination.countTotal"?: boolean;
        "pagination.reverse"?: boolean;
    }, params?: RequestParams) => Promise<HttpResponse<NftQueryDenomsResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryDenom
     * @summary Denom queries the definition of a given denom
     * @request GET:/metachain/nft/denoms/{denomId}
     */
    queryDenom: (denomId: string, params?: RequestParams) => Promise<HttpResponse<NftQueryDenomResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryOwner
     * @summary Owner queries the NFTs of the specified owner
     * @request GET:/metachain/nft/nfts
     */
    queryOwner: (query?: {
        denomId?: string;
        owner?: string;
        "pagination.key"?: string;
        "pagination.offset"?: string;
        "pagination.limit"?: string;
        "pagination.countTotal"?: boolean;
        "pagination.reverse"?: boolean;
    }, params?: RequestParams) => Promise<HttpResponse<NftQueryOwnerResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryNft
     * @summary NFT queries the NFT for the given denom and token ID
     * @request GET:/metachain/nft/nfts/{denomId}/{tokenId}
     */
    queryNft: (denomId: string, tokenId: string, params?: RequestParams) => Promise<HttpResponse<NftQueryNFTResponse, RpcStatus>>;
}
export {};
