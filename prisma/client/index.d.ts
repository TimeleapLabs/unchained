
/**
 * Client
**/

import * as runtime from './runtime/library';
import $Types = runtime.Types // general types
import $Public = runtime.Types.Public
import $Utils = runtime.Types.Utils
import $Extensions = runtime.Types.Extensions
import $Result = runtime.Types.Result

export type PrismaPromise<T> = $Public.PrismaPromise<T>


/**
 * Model AssetPrice
 * 
 */
export type AssetPrice = $Result.DefaultSelection<Prisma.$AssetPricePayload>
/**
 * Model DataSet
 * 
 */
export type DataSet = $Result.DefaultSelection<Prisma.$DataSetPayload>
/**
 * Model Signer
 * 
 */
export type Signer = $Result.DefaultSelection<Prisma.$SignerPayload>
/**
 * Model SignersOnAssetPrice
 * 
 */
export type SignersOnAssetPrice = $Result.DefaultSelection<Prisma.$SignersOnAssetPricePayload>

/**
 * ##  Prisma Client ʲˢ
 * 
 * Type-safe database client for TypeScript & Node.js
 * @example
 * ```
 * const prisma = new PrismaClient()
 * // Fetch zero or more AssetPrices
 * const assetPrices = await prisma.assetPrice.findMany()
 * ```
 *
 * 
 * Read more in our [docs](https://www.prisma.io/docs/reference/tools-and-interfaces/prisma-client).
 */
export class PrismaClient<
  T extends Prisma.PrismaClientOptions = Prisma.PrismaClientOptions,
  U = 'log' extends keyof T ? T['log'] extends Array<Prisma.LogLevel | Prisma.LogDefinition> ? Prisma.GetEvents<T['log']> : never : never,
  ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs
> {
  [K: symbol]: { types: Prisma.TypeMap<ExtArgs>['other'] }

    /**
   * ##  Prisma Client ʲˢ
   * 
   * Type-safe database client for TypeScript & Node.js
   * @example
   * ```
   * const prisma = new PrismaClient()
   * // Fetch zero or more AssetPrices
   * const assetPrices = await prisma.assetPrice.findMany()
   * ```
   *
   * 
   * Read more in our [docs](https://www.prisma.io/docs/reference/tools-and-interfaces/prisma-client).
   */

  constructor(optionsArg ?: Prisma.Subset<T, Prisma.PrismaClientOptions>);
  $on<V extends U>(eventType: V, callback: (event: V extends 'query' ? Prisma.QueryEvent : Prisma.LogEvent) => void): void;

  /**
   * Connect with the database
   */
  $connect(): $Utils.JsPromise<void>;

  /**
   * Disconnect from the database
   */
  $disconnect(): $Utils.JsPromise<void>;

  /**
   * Add a middleware
   * @deprecated since 4.16.0. For new code, prefer client extensions instead.
   * @see https://pris.ly/d/extensions
   */
  $use(cb: Prisma.Middleware): void

/**
   * Executes a prepared raw query and returns the number of affected rows.
   * @example
   * ```
   * const result = await prisma.$executeRaw`UPDATE User SET cool = ${true} WHERE email = ${'user@email.com'};`
   * ```
   * 
   * Read more in our [docs](https://www.prisma.io/docs/reference/tools-and-interfaces/prisma-client/raw-database-access).
   */
  $executeRaw<T = unknown>(query: TemplateStringsArray | Prisma.Sql, ...values: any[]): Prisma.PrismaPromise<number>;

  /**
   * Executes a raw query and returns the number of affected rows.
   * Susceptible to SQL injections, see documentation.
   * @example
   * ```
   * const result = await prisma.$executeRawUnsafe('UPDATE User SET cool = $1 WHERE email = $2 ;', true, 'user@email.com')
   * ```
   * 
   * Read more in our [docs](https://www.prisma.io/docs/reference/tools-and-interfaces/prisma-client/raw-database-access).
   */
  $executeRawUnsafe<T = unknown>(query: string, ...values: any[]): Prisma.PrismaPromise<number>;

  /**
   * Performs a prepared raw query and returns the `SELECT` data.
   * @example
   * ```
   * const result = await prisma.$queryRaw`SELECT * FROM User WHERE id = ${1} OR email = ${'user@email.com'};`
   * ```
   * 
   * Read more in our [docs](https://www.prisma.io/docs/reference/tools-and-interfaces/prisma-client/raw-database-access).
   */
  $queryRaw<T = unknown>(query: TemplateStringsArray | Prisma.Sql, ...values: any[]): Prisma.PrismaPromise<T>;

  /**
   * Performs a raw query and returns the `SELECT` data.
   * Susceptible to SQL injections, see documentation.
   * @example
   * ```
   * const result = await prisma.$queryRawUnsafe('SELECT * FROM User WHERE id = $1 OR email = $2;', 1, 'user@email.com')
   * ```
   * 
   * Read more in our [docs](https://www.prisma.io/docs/reference/tools-and-interfaces/prisma-client/raw-database-access).
   */
  $queryRawUnsafe<T = unknown>(query: string, ...values: any[]): Prisma.PrismaPromise<T>;

  /**
   * Allows the running of a sequence of read/write operations that are guaranteed to either succeed or fail as a whole.
   * @example
   * ```
   * const [george, bob, alice] = await prisma.$transaction([
   *   prisma.user.create({ data: { name: 'George' } }),
   *   prisma.user.create({ data: { name: 'Bob' } }),
   *   prisma.user.create({ data: { name: 'Alice' } }),
   * ])
   * ```
   * 
   * Read more in our [docs](https://www.prisma.io/docs/concepts/components/prisma-client/transactions).
   */
  $transaction<P extends Prisma.PrismaPromise<any>[]>(arg: [...P], options?: { isolationLevel?: Prisma.TransactionIsolationLevel }): $Utils.JsPromise<runtime.Types.Utils.UnwrapTuple<P>>

  $transaction<R>(fn: (prisma: Omit<PrismaClient, runtime.ITXClientDenyList>) => $Utils.JsPromise<R>, options?: { maxWait?: number, timeout?: number, isolationLevel?: Prisma.TransactionIsolationLevel }): $Utils.JsPromise<R>


  $extends: $Extensions.ExtendsHook<'extends', Prisma.TypeMapCb, ExtArgs>

      /**
   * `prisma.assetPrice`: Exposes CRUD operations for the **AssetPrice** model.
    * Example usage:
    * ```ts
    * // Fetch zero or more AssetPrices
    * const assetPrices = await prisma.assetPrice.findMany()
    * ```
    */
  get assetPrice(): Prisma.AssetPriceDelegate<ExtArgs>;

  /**
   * `prisma.dataSet`: Exposes CRUD operations for the **DataSet** model.
    * Example usage:
    * ```ts
    * // Fetch zero or more DataSets
    * const dataSets = await prisma.dataSet.findMany()
    * ```
    */
  get dataSet(): Prisma.DataSetDelegate<ExtArgs>;

  /**
   * `prisma.signer`: Exposes CRUD operations for the **Signer** model.
    * Example usage:
    * ```ts
    * // Fetch zero or more Signers
    * const signers = await prisma.signer.findMany()
    * ```
    */
  get signer(): Prisma.SignerDelegate<ExtArgs>;

  /**
   * `prisma.signersOnAssetPrice`: Exposes CRUD operations for the **SignersOnAssetPrice** model.
    * Example usage:
    * ```ts
    * // Fetch zero or more SignersOnAssetPrices
    * const signersOnAssetPrices = await prisma.signersOnAssetPrice.findMany()
    * ```
    */
  get signersOnAssetPrice(): Prisma.SignersOnAssetPriceDelegate<ExtArgs>;
}

export namespace Prisma {
  export import DMMF = runtime.DMMF

  export type PrismaPromise<T> = $Public.PrismaPromise<T>

  /**
   * Validator
   */
  export import validator = runtime.Public.validator

  /**
   * Prisma Errors
   */
  export import PrismaClientKnownRequestError = runtime.PrismaClientKnownRequestError
  export import PrismaClientUnknownRequestError = runtime.PrismaClientUnknownRequestError
  export import PrismaClientRustPanicError = runtime.PrismaClientRustPanicError
  export import PrismaClientInitializationError = runtime.PrismaClientInitializationError
  export import PrismaClientValidationError = runtime.PrismaClientValidationError
  export import NotFoundError = runtime.NotFoundError

  /**
   * Re-export of sql-template-tag
   */
  export import sql = runtime.sqltag
  export import empty = runtime.empty
  export import join = runtime.join
  export import raw = runtime.raw
  export import Sql = runtime.Sql

  /**
   * Decimal.js
   */
  export import Decimal = runtime.Decimal

  export type DecimalJsLike = runtime.DecimalJsLike

  /**
   * Metrics 
   */
  export type Metrics = runtime.Metrics
  export type Metric<T> = runtime.Metric<T>
  export type MetricHistogram = runtime.MetricHistogram
  export type MetricHistogramBucket = runtime.MetricHistogramBucket

  /**
  * Extensions
  */
  export import Extension = $Extensions.UserArgs
  export import getExtensionContext = runtime.Extensions.getExtensionContext
  export import Args = $Public.Args
  export import Payload = $Public.Payload
  export import Result = $Public.Result
  export import Exact = $Public.Exact

  /**
   * Prisma Client JS version: 5.7.1
   * Query Engine version: 0ca5ccbcfa6bdc81c003cf549abe4269f59c41e5
   */
  export type PrismaVersion = {
    client: string
  }

  export const prismaVersion: PrismaVersion 

  /**
   * Utility Types
   */

  /**
   * From https://github.com/sindresorhus/type-fest/
   * Matches a JSON object.
   * This type can be useful to enforce some input to be JSON-compatible or as a super-type to be extended from. 
   */
  export type JsonObject = {[Key in string]?: JsonValue}

  /**
   * From https://github.com/sindresorhus/type-fest/
   * Matches a JSON array.
   */
  export interface JsonArray extends Array<JsonValue> {}

  /**
   * From https://github.com/sindresorhus/type-fest/
   * Matches any valid JSON value.
   */
  export type JsonValue = string | number | boolean | JsonObject | JsonArray | null

  /**
   * Matches a JSON object.
   * Unlike `JsonObject`, this type allows undefined and read-only properties.
   */
  export type InputJsonObject = {readonly [Key in string]?: InputJsonValue | null}

  /**
   * Matches a JSON array.
   * Unlike `JsonArray`, readonly arrays are assignable to this type.
   */
  export interface InputJsonArray extends ReadonlyArray<InputJsonValue | null> {}

  /**
   * Matches any valid value that can be used as an input for operations like
   * create and update as the value of a JSON field. Unlike `JsonValue`, this
   * type allows read-only arrays and read-only object properties and disallows
   * `null` at the top level.
   *
   * `null` cannot be used as the value of a JSON field because its meaning
   * would be ambiguous. Use `Prisma.JsonNull` to store the JSON null value or
   * `Prisma.DbNull` to clear the JSON value and set the field to the database
   * NULL value instead.
   *
   * @see https://www.prisma.io/docs/concepts/components/prisma-client/working-with-fields/working-with-json-fields#filtering-by-null-values
   */
  export type InputJsonValue = string | number | boolean | InputJsonObject | InputJsonArray | { toJSON(): unknown }

  /**
   * Types of the values used to represent different kinds of `null` values when working with JSON fields.
   * 
   * @see https://www.prisma.io/docs/concepts/components/prisma-client/working-with-fields/working-with-json-fields#filtering-on-a-json-field
   */
  namespace NullTypes {
    /**
    * Type of `Prisma.DbNull`.
    * 
    * You cannot use other instances of this class. Please use the `Prisma.DbNull` value.
    * 
    * @see https://www.prisma.io/docs/concepts/components/prisma-client/working-with-fields/working-with-json-fields#filtering-on-a-json-field
    */
    class DbNull {
      private DbNull: never
      private constructor()
    }

    /**
    * Type of `Prisma.JsonNull`.
    * 
    * You cannot use other instances of this class. Please use the `Prisma.JsonNull` value.
    * 
    * @see https://www.prisma.io/docs/concepts/components/prisma-client/working-with-fields/working-with-json-fields#filtering-on-a-json-field
    */
    class JsonNull {
      private JsonNull: never
      private constructor()
    }

    /**
    * Type of `Prisma.AnyNull`.
    * 
    * You cannot use other instances of this class. Please use the `Prisma.AnyNull` value.
    * 
    * @see https://www.prisma.io/docs/concepts/components/prisma-client/working-with-fields/working-with-json-fields#filtering-on-a-json-field
    */
    class AnyNull {
      private AnyNull: never
      private constructor()
    }
  }

  /**
   * Helper for filtering JSON entries that have `null` on the database (empty on the db)
   * 
   * @see https://www.prisma.io/docs/concepts/components/prisma-client/working-with-fields/working-with-json-fields#filtering-on-a-json-field
   */
  export const DbNull: NullTypes.DbNull

  /**
   * Helper for filtering JSON entries that have JSON `null` values (not empty on the db)
   * 
   * @see https://www.prisma.io/docs/concepts/components/prisma-client/working-with-fields/working-with-json-fields#filtering-on-a-json-field
   */
  export const JsonNull: NullTypes.JsonNull

  /**
   * Helper for filtering JSON entries that are `Prisma.DbNull` or `Prisma.JsonNull`
   * 
   * @see https://www.prisma.io/docs/concepts/components/prisma-client/working-with-fields/working-with-json-fields#filtering-on-a-json-field
   */
  export const AnyNull: NullTypes.AnyNull

  type SelectAndInclude = {
    select: any
    include: any
  }

  /**
   * Get the type of the value, that the Promise holds.
   */
  export type PromiseType<T extends PromiseLike<any>> = T extends PromiseLike<infer U> ? U : T;

  /**
   * Get the return type of a function which returns a Promise.
   */
  export type PromiseReturnType<T extends (...args: any) => $Utils.JsPromise<any>> = PromiseType<ReturnType<T>>

  /**
   * From T, pick a set of properties whose keys are in the union K
   */
  type Prisma__Pick<T, K extends keyof T> = {
      [P in K]: T[P];
  };


  export type Enumerable<T> = T | Array<T>;

  export type RequiredKeys<T> = {
    [K in keyof T]-?: {} extends Prisma__Pick<T, K> ? never : K
  }[keyof T]

  export type TruthyKeys<T> = keyof {
    [K in keyof T as T[K] extends false | undefined | null ? never : K]: K
  }

  export type TrueKeys<T> = TruthyKeys<Prisma__Pick<T, RequiredKeys<T>>>

  /**
   * Subset
   * @desc From `T` pick properties that exist in `U`. Simple version of Intersection
   */
  export type Subset<T, U> = {
    [key in keyof T]: key extends keyof U ? T[key] : never;
  };

  /**
   * SelectSubset
   * @desc From `T` pick properties that exist in `U`. Simple version of Intersection.
   * Additionally, it validates, if both select and include are present. If the case, it errors.
   */
  export type SelectSubset<T, U> = {
    [key in keyof T]: key extends keyof U ? T[key] : never
  } &
    (T extends SelectAndInclude
      ? 'Please either choose `select` or `include`.'
      : {})

  /**
   * Subset + Intersection
   * @desc From `T` pick properties that exist in `U` and intersect `K`
   */
  export type SubsetIntersection<T, U, K> = {
    [key in keyof T]: key extends keyof U ? T[key] : never
  } &
    K

  type Without<T, U> = { [P in Exclude<keyof T, keyof U>]?: never };

  /**
   * XOR is needed to have a real mutually exclusive union type
   * https://stackoverflow.com/questions/42123407/does-typescript-support-mutually-exclusive-types
   */
  type XOR<T, U> =
    T extends object ?
    U extends object ?
      (Without<T, U> & U) | (Without<U, T> & T)
    : U : T


  /**
   * Is T a Record?
   */
  type IsObject<T extends any> = T extends Array<any>
  ? False
  : T extends Date
  ? False
  : T extends Uint8Array
  ? False
  : T extends BigInt
  ? False
  : T extends object
  ? True
  : False


  /**
   * If it's T[], return T
   */
  export type UnEnumerate<T extends unknown> = T extends Array<infer U> ? U : T

  /**
   * From ts-toolbelt
   */

  type __Either<O extends object, K extends Key> = Omit<O, K> &
    {
      // Merge all but K
      [P in K]: Prisma__Pick<O, P & keyof O> // With K possibilities
    }[K]

  type EitherStrict<O extends object, K extends Key> = Strict<__Either<O, K>>

  type EitherLoose<O extends object, K extends Key> = ComputeRaw<__Either<O, K>>

  type _Either<
    O extends object,
    K extends Key,
    strict extends Boolean
  > = {
    1: EitherStrict<O, K>
    0: EitherLoose<O, K>
  }[strict]

  type Either<
    O extends object,
    K extends Key,
    strict extends Boolean = 1
  > = O extends unknown ? _Either<O, K, strict> : never

  export type Union = any

  type PatchUndefined<O extends object, O1 extends object> = {
    [K in keyof O]: O[K] extends undefined ? At<O1, K> : O[K]
  } & {}

  /** Helper Types for "Merge" **/
  export type IntersectOf<U extends Union> = (
    U extends unknown ? (k: U) => void : never
  ) extends (k: infer I) => void
    ? I
    : never

  export type Overwrite<O extends object, O1 extends object> = {
      [K in keyof O]: K extends keyof O1 ? O1[K] : O[K];
  } & {};

  type _Merge<U extends object> = IntersectOf<Overwrite<U, {
      [K in keyof U]-?: At<U, K>;
  }>>;

  type Key = string | number | symbol;
  type AtBasic<O extends object, K extends Key> = K extends keyof O ? O[K] : never;
  type AtStrict<O extends object, K extends Key> = O[K & keyof O];
  type AtLoose<O extends object, K extends Key> = O extends unknown ? AtStrict<O, K> : never;
  export type At<O extends object, K extends Key, strict extends Boolean = 1> = {
      1: AtStrict<O, K>;
      0: AtLoose<O, K>;
  }[strict];

  export type ComputeRaw<A extends any> = A extends Function ? A : {
    [K in keyof A]: A[K];
  } & {};

  export type OptionalFlat<O> = {
    [K in keyof O]?: O[K];
  } & {};

  type _Record<K extends keyof any, T> = {
    [P in K]: T;
  };

  // cause typescript not to expand types and preserve names
  type NoExpand<T> = T extends unknown ? T : never;

  // this type assumes the passed object is entirely optional
  type AtLeast<O extends object, K extends string> = NoExpand<
    O extends unknown
    ? | (K extends keyof O ? { [P in K]: O[P] } & O : O)
      | {[P in keyof O as P extends K ? K : never]-?: O[P]} & O
    : never>;

  type _Strict<U, _U = U> = U extends unknown ? U & OptionalFlat<_Record<Exclude<Keys<_U>, keyof U>, never>> : never;

  export type Strict<U extends object> = ComputeRaw<_Strict<U>>;
  /** End Helper Types for "Merge" **/

  export type Merge<U extends object> = ComputeRaw<_Merge<Strict<U>>>;

  /**
  A [[Boolean]]
  */
  export type Boolean = True | False

  // /**
  // 1
  // */
  export type True = 1

  /**
  0
  */
  export type False = 0

  export type Not<B extends Boolean> = {
    0: 1
    1: 0
  }[B]

  export type Extends<A1 extends any, A2 extends any> = [A1] extends [never]
    ? 0 // anything `never` is false
    : A1 extends A2
    ? 1
    : 0

  export type Has<U extends Union, U1 extends Union> = Not<
    Extends<Exclude<U1, U>, U1>
  >

  export type Or<B1 extends Boolean, B2 extends Boolean> = {
    0: {
      0: 0
      1: 1
    }
    1: {
      0: 1
      1: 1
    }
  }[B1][B2]

  export type Keys<U extends Union> = U extends unknown ? keyof U : never

  type Cast<A, B> = A extends B ? A : B;

  export const type: unique symbol;



  /**
   * Used by group by
   */

  export type GetScalarType<T, O> = O extends object ? {
    [P in keyof T]: P extends keyof O
      ? O[P]
      : never
  } : never

  type FieldPaths<
    T,
    U = Omit<T, '_avg' | '_sum' | '_count' | '_min' | '_max'>
  > = IsObject<T> extends True ? U : T

  type GetHavingFields<T> = {
    [K in keyof T]: Or<
      Or<Extends<'OR', K>, Extends<'AND', K>>,
      Extends<'NOT', K>
    > extends True
      ? // infer is only needed to not hit TS limit
        // based on the brilliant idea of Pierre-Antoine Mills
        // https://github.com/microsoft/TypeScript/issues/30188#issuecomment-478938437
        T[K] extends infer TK
        ? GetHavingFields<UnEnumerate<TK> extends object ? Merge<UnEnumerate<TK>> : never>
        : never
      : {} extends FieldPaths<T[K]>
      ? never
      : K
  }[keyof T]

  /**
   * Convert tuple to union
   */
  type _TupleToUnion<T> = T extends (infer E)[] ? E : never
  type TupleToUnion<K extends readonly any[]> = _TupleToUnion<K>
  type MaybeTupleToUnion<T> = T extends any[] ? TupleToUnion<T> : T

  /**
   * Like `Pick`, but additionally can also accept an array of keys
   */
  type PickEnumerable<T, K extends Enumerable<keyof T> | keyof T> = Prisma__Pick<T, MaybeTupleToUnion<K>>

  /**
   * Exclude all keys with underscores
   */
  type ExcludeUnderscoreKeys<T extends string> = T extends `_${string}` ? never : T


  export type FieldRef<Model, FieldType> = runtime.FieldRef<Model, FieldType>

  type FieldRefInputType<Model, FieldType> = Model extends never ? never : FieldRef<Model, FieldType>


  export const ModelName: {
    AssetPrice: 'AssetPrice',
    DataSet: 'DataSet',
    Signer: 'Signer',
    SignersOnAssetPrice: 'SignersOnAssetPrice'
  };

  export type ModelName = (typeof ModelName)[keyof typeof ModelName]


  export type Datasources = {
    db?: Datasource
  }


  interface TypeMapCb extends $Utils.Fn<{extArgs: $Extensions.InternalArgs}, $Utils.Record<string, any>> {
    returns: Prisma.TypeMap<this['params']['extArgs']>
  }

  export type TypeMap<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    meta: {
      modelProps: 'assetPrice' | 'dataSet' | 'signer' | 'signersOnAssetPrice'
      txIsolationLevel: Prisma.TransactionIsolationLevel
    },
    model: {
      AssetPrice: {
        payload: Prisma.$AssetPricePayload<ExtArgs>
        fields: Prisma.AssetPriceFieldRefs
        operations: {
          findUnique: {
            args: Prisma.AssetPriceFindUniqueArgs<ExtArgs>,
            result: $Utils.PayloadToResult<Prisma.$AssetPricePayload> | null
          }
          findUniqueOrThrow: {
            args: Prisma.AssetPriceFindUniqueOrThrowArgs<ExtArgs>,
            result: $Utils.PayloadToResult<Prisma.$AssetPricePayload>
          }
          findFirst: {
            args: Prisma.AssetPriceFindFirstArgs<ExtArgs>,
            result: $Utils.PayloadToResult<Prisma.$AssetPricePayload> | null
          }
          findFirstOrThrow: {
            args: Prisma.AssetPriceFindFirstOrThrowArgs<ExtArgs>,
            result: $Utils.PayloadToResult<Prisma.$AssetPricePayload>
          }
          findMany: {
            args: Prisma.AssetPriceFindManyArgs<ExtArgs>,
            result: $Utils.PayloadToResult<Prisma.$AssetPricePayload>[]
          }
          create: {
            args: Prisma.AssetPriceCreateArgs<ExtArgs>,
            result: $Utils.PayloadToResult<Prisma.$AssetPricePayload>
          }
          createMany: {
            args: Prisma.AssetPriceCreateManyArgs<ExtArgs>,
            result: Prisma.BatchPayload
          }
          delete: {
            args: Prisma.AssetPriceDeleteArgs<ExtArgs>,
            result: $Utils.PayloadToResult<Prisma.$AssetPricePayload>
          }
          update: {
            args: Prisma.AssetPriceUpdateArgs<ExtArgs>,
            result: $Utils.PayloadToResult<Prisma.$AssetPricePayload>
          }
          deleteMany: {
            args: Prisma.AssetPriceDeleteManyArgs<ExtArgs>,
            result: Prisma.BatchPayload
          }
          updateMany: {
            args: Prisma.AssetPriceUpdateManyArgs<ExtArgs>,
            result: Prisma.BatchPayload
          }
          upsert: {
            args: Prisma.AssetPriceUpsertArgs<ExtArgs>,
            result: $Utils.PayloadToResult<Prisma.$AssetPricePayload>
          }
          aggregate: {
            args: Prisma.AssetPriceAggregateArgs<ExtArgs>,
            result: $Utils.Optional<AggregateAssetPrice>
          }
          groupBy: {
            args: Prisma.AssetPriceGroupByArgs<ExtArgs>,
            result: $Utils.Optional<AssetPriceGroupByOutputType>[]
          }
          count: {
            args: Prisma.AssetPriceCountArgs<ExtArgs>,
            result: $Utils.Optional<AssetPriceCountAggregateOutputType> | number
          }
        }
      }
      DataSet: {
        payload: Prisma.$DataSetPayload<ExtArgs>
        fields: Prisma.DataSetFieldRefs
        operations: {
          findUnique: {
            args: Prisma.DataSetFindUniqueArgs<ExtArgs>,
            result: $Utils.PayloadToResult<Prisma.$DataSetPayload> | null
          }
          findUniqueOrThrow: {
            args: Prisma.DataSetFindUniqueOrThrowArgs<ExtArgs>,
            result: $Utils.PayloadToResult<Prisma.$DataSetPayload>
          }
          findFirst: {
            args: Prisma.DataSetFindFirstArgs<ExtArgs>,
            result: $Utils.PayloadToResult<Prisma.$DataSetPayload> | null
          }
          findFirstOrThrow: {
            args: Prisma.DataSetFindFirstOrThrowArgs<ExtArgs>,
            result: $Utils.PayloadToResult<Prisma.$DataSetPayload>
          }
          findMany: {
            args: Prisma.DataSetFindManyArgs<ExtArgs>,
            result: $Utils.PayloadToResult<Prisma.$DataSetPayload>[]
          }
          create: {
            args: Prisma.DataSetCreateArgs<ExtArgs>,
            result: $Utils.PayloadToResult<Prisma.$DataSetPayload>
          }
          createMany: {
            args: Prisma.DataSetCreateManyArgs<ExtArgs>,
            result: Prisma.BatchPayload
          }
          delete: {
            args: Prisma.DataSetDeleteArgs<ExtArgs>,
            result: $Utils.PayloadToResult<Prisma.$DataSetPayload>
          }
          update: {
            args: Prisma.DataSetUpdateArgs<ExtArgs>,
            result: $Utils.PayloadToResult<Prisma.$DataSetPayload>
          }
          deleteMany: {
            args: Prisma.DataSetDeleteManyArgs<ExtArgs>,
            result: Prisma.BatchPayload
          }
          updateMany: {
            args: Prisma.DataSetUpdateManyArgs<ExtArgs>,
            result: Prisma.BatchPayload
          }
          upsert: {
            args: Prisma.DataSetUpsertArgs<ExtArgs>,
            result: $Utils.PayloadToResult<Prisma.$DataSetPayload>
          }
          aggregate: {
            args: Prisma.DataSetAggregateArgs<ExtArgs>,
            result: $Utils.Optional<AggregateDataSet>
          }
          groupBy: {
            args: Prisma.DataSetGroupByArgs<ExtArgs>,
            result: $Utils.Optional<DataSetGroupByOutputType>[]
          }
          count: {
            args: Prisma.DataSetCountArgs<ExtArgs>,
            result: $Utils.Optional<DataSetCountAggregateOutputType> | number
          }
        }
      }
      Signer: {
        payload: Prisma.$SignerPayload<ExtArgs>
        fields: Prisma.SignerFieldRefs
        operations: {
          findUnique: {
            args: Prisma.SignerFindUniqueArgs<ExtArgs>,
            result: $Utils.PayloadToResult<Prisma.$SignerPayload> | null
          }
          findUniqueOrThrow: {
            args: Prisma.SignerFindUniqueOrThrowArgs<ExtArgs>,
            result: $Utils.PayloadToResult<Prisma.$SignerPayload>
          }
          findFirst: {
            args: Prisma.SignerFindFirstArgs<ExtArgs>,
            result: $Utils.PayloadToResult<Prisma.$SignerPayload> | null
          }
          findFirstOrThrow: {
            args: Prisma.SignerFindFirstOrThrowArgs<ExtArgs>,
            result: $Utils.PayloadToResult<Prisma.$SignerPayload>
          }
          findMany: {
            args: Prisma.SignerFindManyArgs<ExtArgs>,
            result: $Utils.PayloadToResult<Prisma.$SignerPayload>[]
          }
          create: {
            args: Prisma.SignerCreateArgs<ExtArgs>,
            result: $Utils.PayloadToResult<Prisma.$SignerPayload>
          }
          createMany: {
            args: Prisma.SignerCreateManyArgs<ExtArgs>,
            result: Prisma.BatchPayload
          }
          delete: {
            args: Prisma.SignerDeleteArgs<ExtArgs>,
            result: $Utils.PayloadToResult<Prisma.$SignerPayload>
          }
          update: {
            args: Prisma.SignerUpdateArgs<ExtArgs>,
            result: $Utils.PayloadToResult<Prisma.$SignerPayload>
          }
          deleteMany: {
            args: Prisma.SignerDeleteManyArgs<ExtArgs>,
            result: Prisma.BatchPayload
          }
          updateMany: {
            args: Prisma.SignerUpdateManyArgs<ExtArgs>,
            result: Prisma.BatchPayload
          }
          upsert: {
            args: Prisma.SignerUpsertArgs<ExtArgs>,
            result: $Utils.PayloadToResult<Prisma.$SignerPayload>
          }
          aggregate: {
            args: Prisma.SignerAggregateArgs<ExtArgs>,
            result: $Utils.Optional<AggregateSigner>
          }
          groupBy: {
            args: Prisma.SignerGroupByArgs<ExtArgs>,
            result: $Utils.Optional<SignerGroupByOutputType>[]
          }
          count: {
            args: Prisma.SignerCountArgs<ExtArgs>,
            result: $Utils.Optional<SignerCountAggregateOutputType> | number
          }
        }
      }
      SignersOnAssetPrice: {
        payload: Prisma.$SignersOnAssetPricePayload<ExtArgs>
        fields: Prisma.SignersOnAssetPriceFieldRefs
        operations: {
          findUnique: {
            args: Prisma.SignersOnAssetPriceFindUniqueArgs<ExtArgs>,
            result: $Utils.PayloadToResult<Prisma.$SignersOnAssetPricePayload> | null
          }
          findUniqueOrThrow: {
            args: Prisma.SignersOnAssetPriceFindUniqueOrThrowArgs<ExtArgs>,
            result: $Utils.PayloadToResult<Prisma.$SignersOnAssetPricePayload>
          }
          findFirst: {
            args: Prisma.SignersOnAssetPriceFindFirstArgs<ExtArgs>,
            result: $Utils.PayloadToResult<Prisma.$SignersOnAssetPricePayload> | null
          }
          findFirstOrThrow: {
            args: Prisma.SignersOnAssetPriceFindFirstOrThrowArgs<ExtArgs>,
            result: $Utils.PayloadToResult<Prisma.$SignersOnAssetPricePayload>
          }
          findMany: {
            args: Prisma.SignersOnAssetPriceFindManyArgs<ExtArgs>,
            result: $Utils.PayloadToResult<Prisma.$SignersOnAssetPricePayload>[]
          }
          create: {
            args: Prisma.SignersOnAssetPriceCreateArgs<ExtArgs>,
            result: $Utils.PayloadToResult<Prisma.$SignersOnAssetPricePayload>
          }
          createMany: {
            args: Prisma.SignersOnAssetPriceCreateManyArgs<ExtArgs>,
            result: Prisma.BatchPayload
          }
          delete: {
            args: Prisma.SignersOnAssetPriceDeleteArgs<ExtArgs>,
            result: $Utils.PayloadToResult<Prisma.$SignersOnAssetPricePayload>
          }
          update: {
            args: Prisma.SignersOnAssetPriceUpdateArgs<ExtArgs>,
            result: $Utils.PayloadToResult<Prisma.$SignersOnAssetPricePayload>
          }
          deleteMany: {
            args: Prisma.SignersOnAssetPriceDeleteManyArgs<ExtArgs>,
            result: Prisma.BatchPayload
          }
          updateMany: {
            args: Prisma.SignersOnAssetPriceUpdateManyArgs<ExtArgs>,
            result: Prisma.BatchPayload
          }
          upsert: {
            args: Prisma.SignersOnAssetPriceUpsertArgs<ExtArgs>,
            result: $Utils.PayloadToResult<Prisma.$SignersOnAssetPricePayload>
          }
          aggregate: {
            args: Prisma.SignersOnAssetPriceAggregateArgs<ExtArgs>,
            result: $Utils.Optional<AggregateSignersOnAssetPrice>
          }
          groupBy: {
            args: Prisma.SignersOnAssetPriceGroupByArgs<ExtArgs>,
            result: $Utils.Optional<SignersOnAssetPriceGroupByOutputType>[]
          }
          count: {
            args: Prisma.SignersOnAssetPriceCountArgs<ExtArgs>,
            result: $Utils.Optional<SignersOnAssetPriceCountAggregateOutputType> | number
          }
        }
      }
    }
  } & {
    other: {
      payload: any
      operations: {
        $executeRawUnsafe: {
          args: [query: string, ...values: any[]],
          result: any
        }
        $executeRaw: {
          args: [query: TemplateStringsArray | Prisma.Sql, ...values: any[]],
          result: any
        }
        $queryRawUnsafe: {
          args: [query: string, ...values: any[]],
          result: any
        }
        $queryRaw: {
          args: [query: TemplateStringsArray | Prisma.Sql, ...values: any[]],
          result: any
        }
      }
    }
  }
  export const defineExtension: $Extensions.ExtendsHook<'define', Prisma.TypeMapCb, $Extensions.DefaultArgs>
  export type DefaultPrismaClient = PrismaClient
  export type ErrorFormat = 'pretty' | 'colorless' | 'minimal'
  export interface PrismaClientOptions {
    /**
     * Overwrites the datasource url from your schema.prisma file
     */
    datasources?: Datasources
    /**
     * Overwrites the datasource url from your schema.prisma file
     */
    datasourceUrl?: string
    /**
     * @default "colorless"
     */
    errorFormat?: ErrorFormat
    /**
     * @example
     * ```
     * // Defaults to stdout
     * log: ['query', 'info', 'warn', 'error']
     * 
     * // Emit as events
     * log: [
     *   { emit: 'stdout', level: 'query' },
     *   { emit: 'stdout', level: 'info' },
     *   { emit: 'stdout', level: 'warn' }
     *   { emit: 'stdout', level: 'error' }
     * ]
     * ```
     * Read more in our [docs](https://www.prisma.io/docs/reference/tools-and-interfaces/prisma-client/logging#the-log-option).
     */
    log?: (LogLevel | LogDefinition)[]
  }

  /* Types for Logging */
  export type LogLevel = 'info' | 'query' | 'warn' | 'error'
  export type LogDefinition = {
    level: LogLevel
    emit: 'stdout' | 'event'
  }

  export type GetLogType<T extends LogLevel | LogDefinition> = T extends LogDefinition ? T['emit'] extends 'event' ? T['level'] : never : never
  export type GetEvents<T extends any> = T extends Array<LogLevel | LogDefinition> ?
    GetLogType<T[0]> | GetLogType<T[1]> | GetLogType<T[2]> | GetLogType<T[3]>
    : never

  export type QueryEvent = {
    timestamp: Date
    query: string
    params: string
    duration: number
    target: string
  }

  export type LogEvent = {
    timestamp: Date
    message: string
    target: string
  }
  /* End Types for Logging */


  export type PrismaAction =
    | 'findUnique'
    | 'findUniqueOrThrow'
    | 'findMany'
    | 'findFirst'
    | 'findFirstOrThrow'
    | 'create'
    | 'createMany'
    | 'update'
    | 'updateMany'
    | 'upsert'
    | 'delete'
    | 'deleteMany'
    | 'executeRaw'
    | 'queryRaw'
    | 'aggregate'
    | 'count'
    | 'runCommandRaw'
    | 'findRaw'
    | 'groupBy'

  /**
   * These options are being passed into the middleware as "params"
   */
  export type MiddlewareParams = {
    model?: ModelName
    action: PrismaAction
    args: any
    dataPath: string[]
    runInTransaction: boolean
  }

  /**
   * The `T` type makes sure, that the `return proceed` is not forgotten in the middleware implementation
   */
  export type Middleware<T = any> = (
    params: MiddlewareParams,
    next: (params: MiddlewareParams) => $Utils.JsPromise<T>,
  ) => $Utils.JsPromise<T>

  // tested in getLogLevel.test.ts
  export function getLogLevel(log: Array<LogLevel | LogDefinition>): LogLevel | undefined;

  /**
   * `PrismaClient` proxy available in interactive transactions.
   */
  export type TransactionClient = Omit<Prisma.DefaultPrismaClient, runtime.ITXClientDenyList>

  export type Datasource = {
    url?: string
  }

  /**
   * Count Types
   */


  /**
   * Count Type AssetPriceCountOutputType
   */

  export type AssetPriceCountOutputType = {
    signersOnAssetPrice: number
  }

  export type AssetPriceCountOutputTypeSelect<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    signersOnAssetPrice?: boolean | AssetPriceCountOutputTypeCountSignersOnAssetPriceArgs
  }

  // Custom InputTypes

  /**
   * AssetPriceCountOutputType without action
   */
  export type AssetPriceCountOutputTypeDefaultArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    /**
     * Select specific fields to fetch from the AssetPriceCountOutputType
     */
    select?: AssetPriceCountOutputTypeSelect<ExtArgs> | null
  }


  /**
   * AssetPriceCountOutputType without action
   */
  export type AssetPriceCountOutputTypeCountSignersOnAssetPriceArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    where?: SignersOnAssetPriceWhereInput
  }



  /**
   * Count Type DataSetCountOutputType
   */

  export type DataSetCountOutputType = {
    AssetPrice: number
  }

  export type DataSetCountOutputTypeSelect<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    AssetPrice?: boolean | DataSetCountOutputTypeCountAssetPriceArgs
  }

  // Custom InputTypes

  /**
   * DataSetCountOutputType without action
   */
  export type DataSetCountOutputTypeDefaultArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    /**
     * Select specific fields to fetch from the DataSetCountOutputType
     */
    select?: DataSetCountOutputTypeSelect<ExtArgs> | null
  }


  /**
   * DataSetCountOutputType without action
   */
  export type DataSetCountOutputTypeCountAssetPriceArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    where?: AssetPriceWhereInput
  }



  /**
   * Count Type SignerCountOutputType
   */

  export type SignerCountOutputType = {
    signersOnAssetPrice: number
  }

  export type SignerCountOutputTypeSelect<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    signersOnAssetPrice?: boolean | SignerCountOutputTypeCountSignersOnAssetPriceArgs
  }

  // Custom InputTypes

  /**
   * SignerCountOutputType without action
   */
  export type SignerCountOutputTypeDefaultArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    /**
     * Select specific fields to fetch from the SignerCountOutputType
     */
    select?: SignerCountOutputTypeSelect<ExtArgs> | null
  }


  /**
   * SignerCountOutputType without action
   */
  export type SignerCountOutputTypeCountSignersOnAssetPriceArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    where?: SignersOnAssetPriceWhereInput
  }



  /**
   * Models
   */

  /**
   * Model AssetPrice
   */

  export type AggregateAssetPrice = {
    _count: AssetPriceCountAggregateOutputType | null
    _avg: AssetPriceAvgAggregateOutputType | null
    _sum: AssetPriceSumAggregateOutputType | null
    _min: AssetPriceMinAggregateOutputType | null
    _max: AssetPriceMaxAggregateOutputType | null
  }

  export type AssetPriceAvgAggregateOutputType = {
    id: number | null
    dataSetId: number | null
    block: number | null
    price: Decimal | null
  }

  export type AssetPriceSumAggregateOutputType = {
    id: number | null
    dataSetId: number | null
    block: number | null
    price: Decimal | null
  }

  export type AssetPriceMinAggregateOutputType = {
    id: number | null
    dataSetId: number | null
    createdAt: Date | null
    updatedAt: Date | null
    block: number | null
    price: Decimal | null
    signature: string | null
  }

  export type AssetPriceMaxAggregateOutputType = {
    id: number | null
    dataSetId: number | null
    createdAt: Date | null
    updatedAt: Date | null
    block: number | null
    price: Decimal | null
    signature: string | null
  }

  export type AssetPriceCountAggregateOutputType = {
    id: number
    dataSetId: number
    createdAt: number
    updatedAt: number
    block: number
    price: number
    signature: number
    _all: number
  }


  export type AssetPriceAvgAggregateInputType = {
    id?: true
    dataSetId?: true
    block?: true
    price?: true
  }

  export type AssetPriceSumAggregateInputType = {
    id?: true
    dataSetId?: true
    block?: true
    price?: true
  }

  export type AssetPriceMinAggregateInputType = {
    id?: true
    dataSetId?: true
    createdAt?: true
    updatedAt?: true
    block?: true
    price?: true
    signature?: true
  }

  export type AssetPriceMaxAggregateInputType = {
    id?: true
    dataSetId?: true
    createdAt?: true
    updatedAt?: true
    block?: true
    price?: true
    signature?: true
  }

  export type AssetPriceCountAggregateInputType = {
    id?: true
    dataSetId?: true
    createdAt?: true
    updatedAt?: true
    block?: true
    price?: true
    signature?: true
    _all?: true
  }

  export type AssetPriceAggregateArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    /**
     * Filter which AssetPrice to aggregate.
     */
    where?: AssetPriceWhereInput
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/sorting Sorting Docs}
     * 
     * Determine the order of AssetPrices to fetch.
     */
    orderBy?: AssetPriceOrderByWithRelationInput | AssetPriceOrderByWithRelationInput[]
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/pagination#cursor-based-pagination Cursor Docs}
     * 
     * Sets the start position
     */
    cursor?: AssetPriceWhereUniqueInput
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/pagination Pagination Docs}
     * 
     * Take `±n` AssetPrices from the position of the cursor.
     */
    take?: number
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/pagination Pagination Docs}
     * 
     * Skip the first `n` AssetPrices.
     */
    skip?: number
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/aggregations Aggregation Docs}
     * 
     * Count returned AssetPrices
    **/
    _count?: true | AssetPriceCountAggregateInputType
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/aggregations Aggregation Docs}
     * 
     * Select which fields to average
    **/
    _avg?: AssetPriceAvgAggregateInputType
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/aggregations Aggregation Docs}
     * 
     * Select which fields to sum
    **/
    _sum?: AssetPriceSumAggregateInputType
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/aggregations Aggregation Docs}
     * 
     * Select which fields to find the minimum value
    **/
    _min?: AssetPriceMinAggregateInputType
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/aggregations Aggregation Docs}
     * 
     * Select which fields to find the maximum value
    **/
    _max?: AssetPriceMaxAggregateInputType
  }

  export type GetAssetPriceAggregateType<T extends AssetPriceAggregateArgs> = {
        [P in keyof T & keyof AggregateAssetPrice]: P extends '_count' | 'count'
      ? T[P] extends true
        ? number
        : GetScalarType<T[P], AggregateAssetPrice[P]>
      : GetScalarType<T[P], AggregateAssetPrice[P]>
  }




  export type AssetPriceGroupByArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    where?: AssetPriceWhereInput
    orderBy?: AssetPriceOrderByWithAggregationInput | AssetPriceOrderByWithAggregationInput[]
    by: AssetPriceScalarFieldEnum[] | AssetPriceScalarFieldEnum
    having?: AssetPriceScalarWhereWithAggregatesInput
    take?: number
    skip?: number
    _count?: AssetPriceCountAggregateInputType | true
    _avg?: AssetPriceAvgAggregateInputType
    _sum?: AssetPriceSumAggregateInputType
    _min?: AssetPriceMinAggregateInputType
    _max?: AssetPriceMaxAggregateInputType
  }

  export type AssetPriceGroupByOutputType = {
    id: number
    dataSetId: number
    createdAt: Date
    updatedAt: Date
    block: number | null
    price: Decimal
    signature: string
    _count: AssetPriceCountAggregateOutputType | null
    _avg: AssetPriceAvgAggregateOutputType | null
    _sum: AssetPriceSumAggregateOutputType | null
    _min: AssetPriceMinAggregateOutputType | null
    _max: AssetPriceMaxAggregateOutputType | null
  }

  type GetAssetPriceGroupByPayload<T extends AssetPriceGroupByArgs> = Prisma.PrismaPromise<
    Array<
      PickEnumerable<AssetPriceGroupByOutputType, T['by']> &
        {
          [P in ((keyof T) & (keyof AssetPriceGroupByOutputType))]: P extends '_count'
            ? T[P] extends boolean
              ? number
              : GetScalarType<T[P], AssetPriceGroupByOutputType[P]>
            : GetScalarType<T[P], AssetPriceGroupByOutputType[P]>
        }
      >
    >


  export type AssetPriceSelect<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = $Extensions.GetSelect<{
    id?: boolean
    dataSetId?: boolean
    createdAt?: boolean
    updatedAt?: boolean
    block?: boolean
    price?: boolean
    signature?: boolean
    dataset?: boolean | DataSetDefaultArgs<ExtArgs>
    signersOnAssetPrice?: boolean | AssetPrice$signersOnAssetPriceArgs<ExtArgs>
    _count?: boolean | AssetPriceCountOutputTypeDefaultArgs<ExtArgs>
  }, ExtArgs["result"]["assetPrice"]>

  export type AssetPriceSelectScalar = {
    id?: boolean
    dataSetId?: boolean
    createdAt?: boolean
    updatedAt?: boolean
    block?: boolean
    price?: boolean
    signature?: boolean
  }

  export type AssetPriceInclude<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    dataset?: boolean | DataSetDefaultArgs<ExtArgs>
    signersOnAssetPrice?: boolean | AssetPrice$signersOnAssetPriceArgs<ExtArgs>
    _count?: boolean | AssetPriceCountOutputTypeDefaultArgs<ExtArgs>
  }


  export type $AssetPricePayload<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    name: "AssetPrice"
    objects: {
      dataset: Prisma.$DataSetPayload<ExtArgs>
      signersOnAssetPrice: Prisma.$SignersOnAssetPricePayload<ExtArgs>[]
    }
    scalars: $Extensions.GetPayloadResult<{
      id: number
      dataSetId: number
      createdAt: Date
      updatedAt: Date
      block: number | null
      price: Prisma.Decimal
      signature: string
    }, ExtArgs["result"]["assetPrice"]>
    composites: {}
  }


  type AssetPriceGetPayload<S extends boolean | null | undefined | AssetPriceDefaultArgs> = $Result.GetResult<Prisma.$AssetPricePayload, S>

  type AssetPriceCountArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = 
    Omit<AssetPriceFindManyArgs, 'select' | 'include' | 'distinct' > & {
      select?: AssetPriceCountAggregateInputType | true
    }

  export interface AssetPriceDelegate<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> {
    [K: symbol]: { types: Prisma.TypeMap<ExtArgs>['model']['AssetPrice'], meta: { name: 'AssetPrice' } }
    /**
     * Find zero or one AssetPrice that matches the filter.
     * @param {AssetPriceFindUniqueArgs} args - Arguments to find a AssetPrice
     * @example
     * // Get one AssetPrice
     * const assetPrice = await prisma.assetPrice.findUnique({
     *   where: {
     *     // ... provide filter here
     *   }
     * })
    **/
    findUnique<T extends AssetPriceFindUniqueArgs<ExtArgs>>(
      args: SelectSubset<T, AssetPriceFindUniqueArgs<ExtArgs>>
    ): Prisma__AssetPriceClient<$Result.GetResult<Prisma.$AssetPricePayload<ExtArgs>, T, 'findUnique'> | null, null, ExtArgs>

    /**
     * Find one AssetPrice that matches the filter or throw an error  with `error.code='P2025'` 
     *     if no matches were found.
     * @param {AssetPriceFindUniqueOrThrowArgs} args - Arguments to find a AssetPrice
     * @example
     * // Get one AssetPrice
     * const assetPrice = await prisma.assetPrice.findUniqueOrThrow({
     *   where: {
     *     // ... provide filter here
     *   }
     * })
    **/
    findUniqueOrThrow<T extends AssetPriceFindUniqueOrThrowArgs<ExtArgs>>(
      args?: SelectSubset<T, AssetPriceFindUniqueOrThrowArgs<ExtArgs>>
    ): Prisma__AssetPriceClient<$Result.GetResult<Prisma.$AssetPricePayload<ExtArgs>, T, 'findUniqueOrThrow'>, never, ExtArgs>

    /**
     * Find the first AssetPrice that matches the filter.
     * Note, that providing `undefined` is treated as the value not being there.
     * Read more here: https://pris.ly/d/null-undefined
     * @param {AssetPriceFindFirstArgs} args - Arguments to find a AssetPrice
     * @example
     * // Get one AssetPrice
     * const assetPrice = await prisma.assetPrice.findFirst({
     *   where: {
     *     // ... provide filter here
     *   }
     * })
    **/
    findFirst<T extends AssetPriceFindFirstArgs<ExtArgs>>(
      args?: SelectSubset<T, AssetPriceFindFirstArgs<ExtArgs>>
    ): Prisma__AssetPriceClient<$Result.GetResult<Prisma.$AssetPricePayload<ExtArgs>, T, 'findFirst'> | null, null, ExtArgs>

    /**
     * Find the first AssetPrice that matches the filter or
     * throw `PrismaKnownClientError` with `P2025` code if no matches were found.
     * Note, that providing `undefined` is treated as the value not being there.
     * Read more here: https://pris.ly/d/null-undefined
     * @param {AssetPriceFindFirstOrThrowArgs} args - Arguments to find a AssetPrice
     * @example
     * // Get one AssetPrice
     * const assetPrice = await prisma.assetPrice.findFirstOrThrow({
     *   where: {
     *     // ... provide filter here
     *   }
     * })
    **/
    findFirstOrThrow<T extends AssetPriceFindFirstOrThrowArgs<ExtArgs>>(
      args?: SelectSubset<T, AssetPriceFindFirstOrThrowArgs<ExtArgs>>
    ): Prisma__AssetPriceClient<$Result.GetResult<Prisma.$AssetPricePayload<ExtArgs>, T, 'findFirstOrThrow'>, never, ExtArgs>

    /**
     * Find zero or more AssetPrices that matches the filter.
     * Note, that providing `undefined` is treated as the value not being there.
     * Read more here: https://pris.ly/d/null-undefined
     * @param {AssetPriceFindManyArgs=} args - Arguments to filter and select certain fields only.
     * @example
     * // Get all AssetPrices
     * const assetPrices = await prisma.assetPrice.findMany()
     * 
     * // Get first 10 AssetPrices
     * const assetPrices = await prisma.assetPrice.findMany({ take: 10 })
     * 
     * // Only select the `id`
     * const assetPriceWithIdOnly = await prisma.assetPrice.findMany({ select: { id: true } })
     * 
    **/
    findMany<T extends AssetPriceFindManyArgs<ExtArgs>>(
      args?: SelectSubset<T, AssetPriceFindManyArgs<ExtArgs>>
    ): Prisma.PrismaPromise<$Result.GetResult<Prisma.$AssetPricePayload<ExtArgs>, T, 'findMany'>>

    /**
     * Create a AssetPrice.
     * @param {AssetPriceCreateArgs} args - Arguments to create a AssetPrice.
     * @example
     * // Create one AssetPrice
     * const AssetPrice = await prisma.assetPrice.create({
     *   data: {
     *     // ... data to create a AssetPrice
     *   }
     * })
     * 
    **/
    create<T extends AssetPriceCreateArgs<ExtArgs>>(
      args: SelectSubset<T, AssetPriceCreateArgs<ExtArgs>>
    ): Prisma__AssetPriceClient<$Result.GetResult<Prisma.$AssetPricePayload<ExtArgs>, T, 'create'>, never, ExtArgs>

    /**
     * Create many AssetPrices.
     *     @param {AssetPriceCreateManyArgs} args - Arguments to create many AssetPrices.
     *     @example
     *     // Create many AssetPrices
     *     const assetPrice = await prisma.assetPrice.createMany({
     *       data: {
     *         // ... provide data here
     *       }
     *     })
     *     
    **/
    createMany<T extends AssetPriceCreateManyArgs<ExtArgs>>(
      args?: SelectSubset<T, AssetPriceCreateManyArgs<ExtArgs>>
    ): Prisma.PrismaPromise<BatchPayload>

    /**
     * Delete a AssetPrice.
     * @param {AssetPriceDeleteArgs} args - Arguments to delete one AssetPrice.
     * @example
     * // Delete one AssetPrice
     * const AssetPrice = await prisma.assetPrice.delete({
     *   where: {
     *     // ... filter to delete one AssetPrice
     *   }
     * })
     * 
    **/
    delete<T extends AssetPriceDeleteArgs<ExtArgs>>(
      args: SelectSubset<T, AssetPriceDeleteArgs<ExtArgs>>
    ): Prisma__AssetPriceClient<$Result.GetResult<Prisma.$AssetPricePayload<ExtArgs>, T, 'delete'>, never, ExtArgs>

    /**
     * Update one AssetPrice.
     * @param {AssetPriceUpdateArgs} args - Arguments to update one AssetPrice.
     * @example
     * // Update one AssetPrice
     * const assetPrice = await prisma.assetPrice.update({
     *   where: {
     *     // ... provide filter here
     *   },
     *   data: {
     *     // ... provide data here
     *   }
     * })
     * 
    **/
    update<T extends AssetPriceUpdateArgs<ExtArgs>>(
      args: SelectSubset<T, AssetPriceUpdateArgs<ExtArgs>>
    ): Prisma__AssetPriceClient<$Result.GetResult<Prisma.$AssetPricePayload<ExtArgs>, T, 'update'>, never, ExtArgs>

    /**
     * Delete zero or more AssetPrices.
     * @param {AssetPriceDeleteManyArgs} args - Arguments to filter AssetPrices to delete.
     * @example
     * // Delete a few AssetPrices
     * const { count } = await prisma.assetPrice.deleteMany({
     *   where: {
     *     // ... provide filter here
     *   }
     * })
     * 
    **/
    deleteMany<T extends AssetPriceDeleteManyArgs<ExtArgs>>(
      args?: SelectSubset<T, AssetPriceDeleteManyArgs<ExtArgs>>
    ): Prisma.PrismaPromise<BatchPayload>

    /**
     * Update zero or more AssetPrices.
     * Note, that providing `undefined` is treated as the value not being there.
     * Read more here: https://pris.ly/d/null-undefined
     * @param {AssetPriceUpdateManyArgs} args - Arguments to update one or more rows.
     * @example
     * // Update many AssetPrices
     * const assetPrice = await prisma.assetPrice.updateMany({
     *   where: {
     *     // ... provide filter here
     *   },
     *   data: {
     *     // ... provide data here
     *   }
     * })
     * 
    **/
    updateMany<T extends AssetPriceUpdateManyArgs<ExtArgs>>(
      args: SelectSubset<T, AssetPriceUpdateManyArgs<ExtArgs>>
    ): Prisma.PrismaPromise<BatchPayload>

    /**
     * Create or update one AssetPrice.
     * @param {AssetPriceUpsertArgs} args - Arguments to update or create a AssetPrice.
     * @example
     * // Update or create a AssetPrice
     * const assetPrice = await prisma.assetPrice.upsert({
     *   create: {
     *     // ... data to create a AssetPrice
     *   },
     *   update: {
     *     // ... in case it already exists, update
     *   },
     *   where: {
     *     // ... the filter for the AssetPrice we want to update
     *   }
     * })
    **/
    upsert<T extends AssetPriceUpsertArgs<ExtArgs>>(
      args: SelectSubset<T, AssetPriceUpsertArgs<ExtArgs>>
    ): Prisma__AssetPriceClient<$Result.GetResult<Prisma.$AssetPricePayload<ExtArgs>, T, 'upsert'>, never, ExtArgs>

    /**
     * Count the number of AssetPrices.
     * Note, that providing `undefined` is treated as the value not being there.
     * Read more here: https://pris.ly/d/null-undefined
     * @param {AssetPriceCountArgs} args - Arguments to filter AssetPrices to count.
     * @example
     * // Count the number of AssetPrices
     * const count = await prisma.assetPrice.count({
     *   where: {
     *     // ... the filter for the AssetPrices we want to count
     *   }
     * })
    **/
    count<T extends AssetPriceCountArgs>(
      args?: Subset<T, AssetPriceCountArgs>,
    ): Prisma.PrismaPromise<
      T extends $Utils.Record<'select', any>
        ? T['select'] extends true
          ? number
          : GetScalarType<T['select'], AssetPriceCountAggregateOutputType>
        : number
    >

    /**
     * Allows you to perform aggregations operations on a AssetPrice.
     * Note, that providing `undefined` is treated as the value not being there.
     * Read more here: https://pris.ly/d/null-undefined
     * @param {AssetPriceAggregateArgs} args - Select which aggregations you would like to apply and on what fields.
     * @example
     * // Ordered by age ascending
     * // Where email contains prisma.io
     * // Limited to the 10 users
     * const aggregations = await prisma.user.aggregate({
     *   _avg: {
     *     age: true,
     *   },
     *   where: {
     *     email: {
     *       contains: "prisma.io",
     *     },
     *   },
     *   orderBy: {
     *     age: "asc",
     *   },
     *   take: 10,
     * })
    **/
    aggregate<T extends AssetPriceAggregateArgs>(args: Subset<T, AssetPriceAggregateArgs>): Prisma.PrismaPromise<GetAssetPriceAggregateType<T>>

    /**
     * Group by AssetPrice.
     * Note, that providing `undefined` is treated as the value not being there.
     * Read more here: https://pris.ly/d/null-undefined
     * @param {AssetPriceGroupByArgs} args - Group by arguments.
     * @example
     * // Group by city, order by createdAt, get count
     * const result = await prisma.user.groupBy({
     *   by: ['city', 'createdAt'],
     *   orderBy: {
     *     createdAt: true
     *   },
     *   _count: {
     *     _all: true
     *   },
     * })
     * 
    **/
    groupBy<
      T extends AssetPriceGroupByArgs,
      HasSelectOrTake extends Or<
        Extends<'skip', Keys<T>>,
        Extends<'take', Keys<T>>
      >,
      OrderByArg extends True extends HasSelectOrTake
        ? { orderBy: AssetPriceGroupByArgs['orderBy'] }
        : { orderBy?: AssetPriceGroupByArgs['orderBy'] },
      OrderFields extends ExcludeUnderscoreKeys<Keys<MaybeTupleToUnion<T['orderBy']>>>,
      ByFields extends MaybeTupleToUnion<T['by']>,
      ByValid extends Has<ByFields, OrderFields>,
      HavingFields extends GetHavingFields<T['having']>,
      HavingValid extends Has<ByFields, HavingFields>,
      ByEmpty extends T['by'] extends never[] ? True : False,
      InputErrors extends ByEmpty extends True
      ? `Error: "by" must not be empty.`
      : HavingValid extends False
      ? {
          [P in HavingFields]: P extends ByFields
            ? never
            : P extends string
            ? `Error: Field "${P}" used in "having" needs to be provided in "by".`
            : [
                Error,
                'Field ',
                P,
                ` in "having" needs to be provided in "by"`,
              ]
        }[HavingFields]
      : 'take' extends Keys<T>
      ? 'orderBy' extends Keys<T>
        ? ByValid extends True
          ? {}
          : {
              [P in OrderFields]: P extends ByFields
                ? never
                : `Error: Field "${P}" in "orderBy" needs to be provided in "by"`
            }[OrderFields]
        : 'Error: If you provide "take", you also need to provide "orderBy"'
      : 'skip' extends Keys<T>
      ? 'orderBy' extends Keys<T>
        ? ByValid extends True
          ? {}
          : {
              [P in OrderFields]: P extends ByFields
                ? never
                : `Error: Field "${P}" in "orderBy" needs to be provided in "by"`
            }[OrderFields]
        : 'Error: If you provide "skip", you also need to provide "orderBy"'
      : ByValid extends True
      ? {}
      : {
          [P in OrderFields]: P extends ByFields
            ? never
            : `Error: Field "${P}" in "orderBy" needs to be provided in "by"`
        }[OrderFields]
    >(args: SubsetIntersection<T, AssetPriceGroupByArgs, OrderByArg> & InputErrors): {} extends InputErrors ? GetAssetPriceGroupByPayload<T> : Prisma.PrismaPromise<InputErrors>
  /**
   * Fields of the AssetPrice model
   */
  readonly fields: AssetPriceFieldRefs;
  }

  /**
   * The delegate class that acts as a "Promise-like" for AssetPrice.
   * Why is this prefixed with `Prisma__`?
   * Because we want to prevent naming conflicts as mentioned in
   * https://github.com/prisma/prisma-client-js/issues/707
   */
  export interface Prisma__AssetPriceClient<T, Null = never, ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> extends Prisma.PrismaPromise<T> {
    readonly [Symbol.toStringTag]: 'PrismaPromise';

    dataset<T extends DataSetDefaultArgs<ExtArgs> = {}>(args?: Subset<T, DataSetDefaultArgs<ExtArgs>>): Prisma__DataSetClient<$Result.GetResult<Prisma.$DataSetPayload<ExtArgs>, T, 'findUniqueOrThrow'> | Null, Null, ExtArgs>;

    signersOnAssetPrice<T extends AssetPrice$signersOnAssetPriceArgs<ExtArgs> = {}>(args?: Subset<T, AssetPrice$signersOnAssetPriceArgs<ExtArgs>>): Prisma.PrismaPromise<$Result.GetResult<Prisma.$SignersOnAssetPricePayload<ExtArgs>, T, 'findMany'> | Null>;

    /**
     * Attaches callbacks for the resolution and/or rejection of the Promise.
     * @param onfulfilled The callback to execute when the Promise is resolved.
     * @param onrejected The callback to execute when the Promise is rejected.
     * @returns A Promise for the completion of which ever callback is executed.
     */
    then<TResult1 = T, TResult2 = never>(onfulfilled?: ((value: T) => TResult1 | PromiseLike<TResult1>) | undefined | null, onrejected?: ((reason: any) => TResult2 | PromiseLike<TResult2>) | undefined | null): $Utils.JsPromise<TResult1 | TResult2>;
    /**
     * Attaches a callback for only the rejection of the Promise.
     * @param onrejected The callback to execute when the Promise is rejected.
     * @returns A Promise for the completion of the callback.
     */
    catch<TResult = never>(onrejected?: ((reason: any) => TResult | PromiseLike<TResult>) | undefined | null): $Utils.JsPromise<T | TResult>;
    /**
     * Attaches a callback that is invoked when the Promise is settled (fulfilled or rejected). The
     * resolved value cannot be modified from the callback.
     * @param onfinally The callback to execute when the Promise is settled (fulfilled or rejected).
     * @returns A Promise for the completion of the callback.
     */
    finally(onfinally?: (() => void) | undefined | null): $Utils.JsPromise<T>;
  }



  /**
   * Fields of the AssetPrice model
   */ 
  interface AssetPriceFieldRefs {
    readonly id: FieldRef<"AssetPrice", 'Int'>
    readonly dataSetId: FieldRef<"AssetPrice", 'Int'>
    readonly createdAt: FieldRef<"AssetPrice", 'DateTime'>
    readonly updatedAt: FieldRef<"AssetPrice", 'DateTime'>
    readonly block: FieldRef<"AssetPrice", 'Int'>
    readonly price: FieldRef<"AssetPrice", 'Decimal'>
    readonly signature: FieldRef<"AssetPrice", 'String'>
  }
    

  // Custom InputTypes

  /**
   * AssetPrice findUnique
   */
  export type AssetPriceFindUniqueArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    /**
     * Select specific fields to fetch from the AssetPrice
     */
    select?: AssetPriceSelect<ExtArgs> | null
    /**
     * Choose, which related nodes to fetch as well.
     */
    include?: AssetPriceInclude<ExtArgs> | null
    /**
     * Filter, which AssetPrice to fetch.
     */
    where: AssetPriceWhereUniqueInput
  }


  /**
   * AssetPrice findUniqueOrThrow
   */
  export type AssetPriceFindUniqueOrThrowArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    /**
     * Select specific fields to fetch from the AssetPrice
     */
    select?: AssetPriceSelect<ExtArgs> | null
    /**
     * Choose, which related nodes to fetch as well.
     */
    include?: AssetPriceInclude<ExtArgs> | null
    /**
     * Filter, which AssetPrice to fetch.
     */
    where: AssetPriceWhereUniqueInput
  }


  /**
   * AssetPrice findFirst
   */
  export type AssetPriceFindFirstArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    /**
     * Select specific fields to fetch from the AssetPrice
     */
    select?: AssetPriceSelect<ExtArgs> | null
    /**
     * Choose, which related nodes to fetch as well.
     */
    include?: AssetPriceInclude<ExtArgs> | null
    /**
     * Filter, which AssetPrice to fetch.
     */
    where?: AssetPriceWhereInput
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/sorting Sorting Docs}
     * 
     * Determine the order of AssetPrices to fetch.
     */
    orderBy?: AssetPriceOrderByWithRelationInput | AssetPriceOrderByWithRelationInput[]
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/pagination#cursor-based-pagination Cursor Docs}
     * 
     * Sets the position for searching for AssetPrices.
     */
    cursor?: AssetPriceWhereUniqueInput
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/pagination Pagination Docs}
     * 
     * Take `±n` AssetPrices from the position of the cursor.
     */
    take?: number
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/pagination Pagination Docs}
     * 
     * Skip the first `n` AssetPrices.
     */
    skip?: number
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/distinct Distinct Docs}
     * 
     * Filter by unique combinations of AssetPrices.
     */
    distinct?: AssetPriceScalarFieldEnum | AssetPriceScalarFieldEnum[]
  }


  /**
   * AssetPrice findFirstOrThrow
   */
  export type AssetPriceFindFirstOrThrowArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    /**
     * Select specific fields to fetch from the AssetPrice
     */
    select?: AssetPriceSelect<ExtArgs> | null
    /**
     * Choose, which related nodes to fetch as well.
     */
    include?: AssetPriceInclude<ExtArgs> | null
    /**
     * Filter, which AssetPrice to fetch.
     */
    where?: AssetPriceWhereInput
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/sorting Sorting Docs}
     * 
     * Determine the order of AssetPrices to fetch.
     */
    orderBy?: AssetPriceOrderByWithRelationInput | AssetPriceOrderByWithRelationInput[]
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/pagination#cursor-based-pagination Cursor Docs}
     * 
     * Sets the position for searching for AssetPrices.
     */
    cursor?: AssetPriceWhereUniqueInput
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/pagination Pagination Docs}
     * 
     * Take `±n` AssetPrices from the position of the cursor.
     */
    take?: number
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/pagination Pagination Docs}
     * 
     * Skip the first `n` AssetPrices.
     */
    skip?: number
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/distinct Distinct Docs}
     * 
     * Filter by unique combinations of AssetPrices.
     */
    distinct?: AssetPriceScalarFieldEnum | AssetPriceScalarFieldEnum[]
  }


  /**
   * AssetPrice findMany
   */
  export type AssetPriceFindManyArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    /**
     * Select specific fields to fetch from the AssetPrice
     */
    select?: AssetPriceSelect<ExtArgs> | null
    /**
     * Choose, which related nodes to fetch as well.
     */
    include?: AssetPriceInclude<ExtArgs> | null
    /**
     * Filter, which AssetPrices to fetch.
     */
    where?: AssetPriceWhereInput
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/sorting Sorting Docs}
     * 
     * Determine the order of AssetPrices to fetch.
     */
    orderBy?: AssetPriceOrderByWithRelationInput | AssetPriceOrderByWithRelationInput[]
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/pagination#cursor-based-pagination Cursor Docs}
     * 
     * Sets the position for listing AssetPrices.
     */
    cursor?: AssetPriceWhereUniqueInput
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/pagination Pagination Docs}
     * 
     * Take `±n` AssetPrices from the position of the cursor.
     */
    take?: number
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/pagination Pagination Docs}
     * 
     * Skip the first `n` AssetPrices.
     */
    skip?: number
    distinct?: AssetPriceScalarFieldEnum | AssetPriceScalarFieldEnum[]
  }


  /**
   * AssetPrice create
   */
  export type AssetPriceCreateArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    /**
     * Select specific fields to fetch from the AssetPrice
     */
    select?: AssetPriceSelect<ExtArgs> | null
    /**
     * Choose, which related nodes to fetch as well.
     */
    include?: AssetPriceInclude<ExtArgs> | null
    /**
     * The data needed to create a AssetPrice.
     */
    data: XOR<AssetPriceCreateInput, AssetPriceUncheckedCreateInput>
  }


  /**
   * AssetPrice createMany
   */
  export type AssetPriceCreateManyArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    /**
     * The data used to create many AssetPrices.
     */
    data: AssetPriceCreateManyInput | AssetPriceCreateManyInput[]
    skipDuplicates?: boolean
  }


  /**
   * AssetPrice update
   */
  export type AssetPriceUpdateArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    /**
     * Select specific fields to fetch from the AssetPrice
     */
    select?: AssetPriceSelect<ExtArgs> | null
    /**
     * Choose, which related nodes to fetch as well.
     */
    include?: AssetPriceInclude<ExtArgs> | null
    /**
     * The data needed to update a AssetPrice.
     */
    data: XOR<AssetPriceUpdateInput, AssetPriceUncheckedUpdateInput>
    /**
     * Choose, which AssetPrice to update.
     */
    where: AssetPriceWhereUniqueInput
  }


  /**
   * AssetPrice updateMany
   */
  export type AssetPriceUpdateManyArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    /**
     * The data used to update AssetPrices.
     */
    data: XOR<AssetPriceUpdateManyMutationInput, AssetPriceUncheckedUpdateManyInput>
    /**
     * Filter which AssetPrices to update
     */
    where?: AssetPriceWhereInput
  }


  /**
   * AssetPrice upsert
   */
  export type AssetPriceUpsertArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    /**
     * Select specific fields to fetch from the AssetPrice
     */
    select?: AssetPriceSelect<ExtArgs> | null
    /**
     * Choose, which related nodes to fetch as well.
     */
    include?: AssetPriceInclude<ExtArgs> | null
    /**
     * The filter to search for the AssetPrice to update in case it exists.
     */
    where: AssetPriceWhereUniqueInput
    /**
     * In case the AssetPrice found by the `where` argument doesn't exist, create a new AssetPrice with this data.
     */
    create: XOR<AssetPriceCreateInput, AssetPriceUncheckedCreateInput>
    /**
     * In case the AssetPrice was found with the provided `where` argument, update it with this data.
     */
    update: XOR<AssetPriceUpdateInput, AssetPriceUncheckedUpdateInput>
  }


  /**
   * AssetPrice delete
   */
  export type AssetPriceDeleteArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    /**
     * Select specific fields to fetch from the AssetPrice
     */
    select?: AssetPriceSelect<ExtArgs> | null
    /**
     * Choose, which related nodes to fetch as well.
     */
    include?: AssetPriceInclude<ExtArgs> | null
    /**
     * Filter which AssetPrice to delete.
     */
    where: AssetPriceWhereUniqueInput
  }


  /**
   * AssetPrice deleteMany
   */
  export type AssetPriceDeleteManyArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    /**
     * Filter which AssetPrices to delete
     */
    where?: AssetPriceWhereInput
  }


  /**
   * AssetPrice.signersOnAssetPrice
   */
  export type AssetPrice$signersOnAssetPriceArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    /**
     * Select specific fields to fetch from the SignersOnAssetPrice
     */
    select?: SignersOnAssetPriceSelect<ExtArgs> | null
    /**
     * Choose, which related nodes to fetch as well.
     */
    include?: SignersOnAssetPriceInclude<ExtArgs> | null
    where?: SignersOnAssetPriceWhereInput
    orderBy?: SignersOnAssetPriceOrderByWithRelationInput | SignersOnAssetPriceOrderByWithRelationInput[]
    cursor?: SignersOnAssetPriceWhereUniqueInput
    take?: number
    skip?: number
    distinct?: SignersOnAssetPriceScalarFieldEnum | SignersOnAssetPriceScalarFieldEnum[]
  }


  /**
   * AssetPrice without action
   */
  export type AssetPriceDefaultArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    /**
     * Select specific fields to fetch from the AssetPrice
     */
    select?: AssetPriceSelect<ExtArgs> | null
    /**
     * Choose, which related nodes to fetch as well.
     */
    include?: AssetPriceInclude<ExtArgs> | null
  }



  /**
   * Model DataSet
   */

  export type AggregateDataSet = {
    _count: DataSetCountAggregateOutputType | null
    _avg: DataSetAvgAggregateOutputType | null
    _sum: DataSetSumAggregateOutputType | null
    _min: DataSetMinAggregateOutputType | null
    _max: DataSetMaxAggregateOutputType | null
  }

  export type DataSetAvgAggregateOutputType = {
    id: number | null
  }

  export type DataSetSumAggregateOutputType = {
    id: number | null
  }

  export type DataSetMinAggregateOutputType = {
    id: number | null
    name: string | null
  }

  export type DataSetMaxAggregateOutputType = {
    id: number | null
    name: string | null
  }

  export type DataSetCountAggregateOutputType = {
    id: number
    name: number
    _all: number
  }


  export type DataSetAvgAggregateInputType = {
    id?: true
  }

  export type DataSetSumAggregateInputType = {
    id?: true
  }

  export type DataSetMinAggregateInputType = {
    id?: true
    name?: true
  }

  export type DataSetMaxAggregateInputType = {
    id?: true
    name?: true
  }

  export type DataSetCountAggregateInputType = {
    id?: true
    name?: true
    _all?: true
  }

  export type DataSetAggregateArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    /**
     * Filter which DataSet to aggregate.
     */
    where?: DataSetWhereInput
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/sorting Sorting Docs}
     * 
     * Determine the order of DataSets to fetch.
     */
    orderBy?: DataSetOrderByWithRelationInput | DataSetOrderByWithRelationInput[]
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/pagination#cursor-based-pagination Cursor Docs}
     * 
     * Sets the start position
     */
    cursor?: DataSetWhereUniqueInput
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/pagination Pagination Docs}
     * 
     * Take `±n` DataSets from the position of the cursor.
     */
    take?: number
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/pagination Pagination Docs}
     * 
     * Skip the first `n` DataSets.
     */
    skip?: number
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/aggregations Aggregation Docs}
     * 
     * Count returned DataSets
    **/
    _count?: true | DataSetCountAggregateInputType
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/aggregations Aggregation Docs}
     * 
     * Select which fields to average
    **/
    _avg?: DataSetAvgAggregateInputType
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/aggregations Aggregation Docs}
     * 
     * Select which fields to sum
    **/
    _sum?: DataSetSumAggregateInputType
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/aggregations Aggregation Docs}
     * 
     * Select which fields to find the minimum value
    **/
    _min?: DataSetMinAggregateInputType
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/aggregations Aggregation Docs}
     * 
     * Select which fields to find the maximum value
    **/
    _max?: DataSetMaxAggregateInputType
  }

  export type GetDataSetAggregateType<T extends DataSetAggregateArgs> = {
        [P in keyof T & keyof AggregateDataSet]: P extends '_count' | 'count'
      ? T[P] extends true
        ? number
        : GetScalarType<T[P], AggregateDataSet[P]>
      : GetScalarType<T[P], AggregateDataSet[P]>
  }




  export type DataSetGroupByArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    where?: DataSetWhereInput
    orderBy?: DataSetOrderByWithAggregationInput | DataSetOrderByWithAggregationInput[]
    by: DataSetScalarFieldEnum[] | DataSetScalarFieldEnum
    having?: DataSetScalarWhereWithAggregatesInput
    take?: number
    skip?: number
    _count?: DataSetCountAggregateInputType | true
    _avg?: DataSetAvgAggregateInputType
    _sum?: DataSetSumAggregateInputType
    _min?: DataSetMinAggregateInputType
    _max?: DataSetMaxAggregateInputType
  }

  export type DataSetGroupByOutputType = {
    id: number
    name: string
    _count: DataSetCountAggregateOutputType | null
    _avg: DataSetAvgAggregateOutputType | null
    _sum: DataSetSumAggregateOutputType | null
    _min: DataSetMinAggregateOutputType | null
    _max: DataSetMaxAggregateOutputType | null
  }

  type GetDataSetGroupByPayload<T extends DataSetGroupByArgs> = Prisma.PrismaPromise<
    Array<
      PickEnumerable<DataSetGroupByOutputType, T['by']> &
        {
          [P in ((keyof T) & (keyof DataSetGroupByOutputType))]: P extends '_count'
            ? T[P] extends boolean
              ? number
              : GetScalarType<T[P], DataSetGroupByOutputType[P]>
            : GetScalarType<T[P], DataSetGroupByOutputType[P]>
        }
      >
    >


  export type DataSetSelect<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = $Extensions.GetSelect<{
    id?: boolean
    name?: boolean
    AssetPrice?: boolean | DataSet$AssetPriceArgs<ExtArgs>
    _count?: boolean | DataSetCountOutputTypeDefaultArgs<ExtArgs>
  }, ExtArgs["result"]["dataSet"]>

  export type DataSetSelectScalar = {
    id?: boolean
    name?: boolean
  }

  export type DataSetInclude<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    AssetPrice?: boolean | DataSet$AssetPriceArgs<ExtArgs>
    _count?: boolean | DataSetCountOutputTypeDefaultArgs<ExtArgs>
  }


  export type $DataSetPayload<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    name: "DataSet"
    objects: {
      AssetPrice: Prisma.$AssetPricePayload<ExtArgs>[]
    }
    scalars: $Extensions.GetPayloadResult<{
      id: number
      name: string
    }, ExtArgs["result"]["dataSet"]>
    composites: {}
  }


  type DataSetGetPayload<S extends boolean | null | undefined | DataSetDefaultArgs> = $Result.GetResult<Prisma.$DataSetPayload, S>

  type DataSetCountArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = 
    Omit<DataSetFindManyArgs, 'select' | 'include' | 'distinct' > & {
      select?: DataSetCountAggregateInputType | true
    }

  export interface DataSetDelegate<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> {
    [K: symbol]: { types: Prisma.TypeMap<ExtArgs>['model']['DataSet'], meta: { name: 'DataSet' } }
    /**
     * Find zero or one DataSet that matches the filter.
     * @param {DataSetFindUniqueArgs} args - Arguments to find a DataSet
     * @example
     * // Get one DataSet
     * const dataSet = await prisma.dataSet.findUnique({
     *   where: {
     *     // ... provide filter here
     *   }
     * })
    **/
    findUnique<T extends DataSetFindUniqueArgs<ExtArgs>>(
      args: SelectSubset<T, DataSetFindUniqueArgs<ExtArgs>>
    ): Prisma__DataSetClient<$Result.GetResult<Prisma.$DataSetPayload<ExtArgs>, T, 'findUnique'> | null, null, ExtArgs>

    /**
     * Find one DataSet that matches the filter or throw an error  with `error.code='P2025'` 
     *     if no matches were found.
     * @param {DataSetFindUniqueOrThrowArgs} args - Arguments to find a DataSet
     * @example
     * // Get one DataSet
     * const dataSet = await prisma.dataSet.findUniqueOrThrow({
     *   where: {
     *     // ... provide filter here
     *   }
     * })
    **/
    findUniqueOrThrow<T extends DataSetFindUniqueOrThrowArgs<ExtArgs>>(
      args?: SelectSubset<T, DataSetFindUniqueOrThrowArgs<ExtArgs>>
    ): Prisma__DataSetClient<$Result.GetResult<Prisma.$DataSetPayload<ExtArgs>, T, 'findUniqueOrThrow'>, never, ExtArgs>

    /**
     * Find the first DataSet that matches the filter.
     * Note, that providing `undefined` is treated as the value not being there.
     * Read more here: https://pris.ly/d/null-undefined
     * @param {DataSetFindFirstArgs} args - Arguments to find a DataSet
     * @example
     * // Get one DataSet
     * const dataSet = await prisma.dataSet.findFirst({
     *   where: {
     *     // ... provide filter here
     *   }
     * })
    **/
    findFirst<T extends DataSetFindFirstArgs<ExtArgs>>(
      args?: SelectSubset<T, DataSetFindFirstArgs<ExtArgs>>
    ): Prisma__DataSetClient<$Result.GetResult<Prisma.$DataSetPayload<ExtArgs>, T, 'findFirst'> | null, null, ExtArgs>

    /**
     * Find the first DataSet that matches the filter or
     * throw `PrismaKnownClientError` with `P2025` code if no matches were found.
     * Note, that providing `undefined` is treated as the value not being there.
     * Read more here: https://pris.ly/d/null-undefined
     * @param {DataSetFindFirstOrThrowArgs} args - Arguments to find a DataSet
     * @example
     * // Get one DataSet
     * const dataSet = await prisma.dataSet.findFirstOrThrow({
     *   where: {
     *     // ... provide filter here
     *   }
     * })
    **/
    findFirstOrThrow<T extends DataSetFindFirstOrThrowArgs<ExtArgs>>(
      args?: SelectSubset<T, DataSetFindFirstOrThrowArgs<ExtArgs>>
    ): Prisma__DataSetClient<$Result.GetResult<Prisma.$DataSetPayload<ExtArgs>, T, 'findFirstOrThrow'>, never, ExtArgs>

    /**
     * Find zero or more DataSets that matches the filter.
     * Note, that providing `undefined` is treated as the value not being there.
     * Read more here: https://pris.ly/d/null-undefined
     * @param {DataSetFindManyArgs=} args - Arguments to filter and select certain fields only.
     * @example
     * // Get all DataSets
     * const dataSets = await prisma.dataSet.findMany()
     * 
     * // Get first 10 DataSets
     * const dataSets = await prisma.dataSet.findMany({ take: 10 })
     * 
     * // Only select the `id`
     * const dataSetWithIdOnly = await prisma.dataSet.findMany({ select: { id: true } })
     * 
    **/
    findMany<T extends DataSetFindManyArgs<ExtArgs>>(
      args?: SelectSubset<T, DataSetFindManyArgs<ExtArgs>>
    ): Prisma.PrismaPromise<$Result.GetResult<Prisma.$DataSetPayload<ExtArgs>, T, 'findMany'>>

    /**
     * Create a DataSet.
     * @param {DataSetCreateArgs} args - Arguments to create a DataSet.
     * @example
     * // Create one DataSet
     * const DataSet = await prisma.dataSet.create({
     *   data: {
     *     // ... data to create a DataSet
     *   }
     * })
     * 
    **/
    create<T extends DataSetCreateArgs<ExtArgs>>(
      args: SelectSubset<T, DataSetCreateArgs<ExtArgs>>
    ): Prisma__DataSetClient<$Result.GetResult<Prisma.$DataSetPayload<ExtArgs>, T, 'create'>, never, ExtArgs>

    /**
     * Create many DataSets.
     *     @param {DataSetCreateManyArgs} args - Arguments to create many DataSets.
     *     @example
     *     // Create many DataSets
     *     const dataSet = await prisma.dataSet.createMany({
     *       data: {
     *         // ... provide data here
     *       }
     *     })
     *     
    **/
    createMany<T extends DataSetCreateManyArgs<ExtArgs>>(
      args?: SelectSubset<T, DataSetCreateManyArgs<ExtArgs>>
    ): Prisma.PrismaPromise<BatchPayload>

    /**
     * Delete a DataSet.
     * @param {DataSetDeleteArgs} args - Arguments to delete one DataSet.
     * @example
     * // Delete one DataSet
     * const DataSet = await prisma.dataSet.delete({
     *   where: {
     *     // ... filter to delete one DataSet
     *   }
     * })
     * 
    **/
    delete<T extends DataSetDeleteArgs<ExtArgs>>(
      args: SelectSubset<T, DataSetDeleteArgs<ExtArgs>>
    ): Prisma__DataSetClient<$Result.GetResult<Prisma.$DataSetPayload<ExtArgs>, T, 'delete'>, never, ExtArgs>

    /**
     * Update one DataSet.
     * @param {DataSetUpdateArgs} args - Arguments to update one DataSet.
     * @example
     * // Update one DataSet
     * const dataSet = await prisma.dataSet.update({
     *   where: {
     *     // ... provide filter here
     *   },
     *   data: {
     *     // ... provide data here
     *   }
     * })
     * 
    **/
    update<T extends DataSetUpdateArgs<ExtArgs>>(
      args: SelectSubset<T, DataSetUpdateArgs<ExtArgs>>
    ): Prisma__DataSetClient<$Result.GetResult<Prisma.$DataSetPayload<ExtArgs>, T, 'update'>, never, ExtArgs>

    /**
     * Delete zero or more DataSets.
     * @param {DataSetDeleteManyArgs} args - Arguments to filter DataSets to delete.
     * @example
     * // Delete a few DataSets
     * const { count } = await prisma.dataSet.deleteMany({
     *   where: {
     *     // ... provide filter here
     *   }
     * })
     * 
    **/
    deleteMany<T extends DataSetDeleteManyArgs<ExtArgs>>(
      args?: SelectSubset<T, DataSetDeleteManyArgs<ExtArgs>>
    ): Prisma.PrismaPromise<BatchPayload>

    /**
     * Update zero or more DataSets.
     * Note, that providing `undefined` is treated as the value not being there.
     * Read more here: https://pris.ly/d/null-undefined
     * @param {DataSetUpdateManyArgs} args - Arguments to update one or more rows.
     * @example
     * // Update many DataSets
     * const dataSet = await prisma.dataSet.updateMany({
     *   where: {
     *     // ... provide filter here
     *   },
     *   data: {
     *     // ... provide data here
     *   }
     * })
     * 
    **/
    updateMany<T extends DataSetUpdateManyArgs<ExtArgs>>(
      args: SelectSubset<T, DataSetUpdateManyArgs<ExtArgs>>
    ): Prisma.PrismaPromise<BatchPayload>

    /**
     * Create or update one DataSet.
     * @param {DataSetUpsertArgs} args - Arguments to update or create a DataSet.
     * @example
     * // Update or create a DataSet
     * const dataSet = await prisma.dataSet.upsert({
     *   create: {
     *     // ... data to create a DataSet
     *   },
     *   update: {
     *     // ... in case it already exists, update
     *   },
     *   where: {
     *     // ... the filter for the DataSet we want to update
     *   }
     * })
    **/
    upsert<T extends DataSetUpsertArgs<ExtArgs>>(
      args: SelectSubset<T, DataSetUpsertArgs<ExtArgs>>
    ): Prisma__DataSetClient<$Result.GetResult<Prisma.$DataSetPayload<ExtArgs>, T, 'upsert'>, never, ExtArgs>

    /**
     * Count the number of DataSets.
     * Note, that providing `undefined` is treated as the value not being there.
     * Read more here: https://pris.ly/d/null-undefined
     * @param {DataSetCountArgs} args - Arguments to filter DataSets to count.
     * @example
     * // Count the number of DataSets
     * const count = await prisma.dataSet.count({
     *   where: {
     *     // ... the filter for the DataSets we want to count
     *   }
     * })
    **/
    count<T extends DataSetCountArgs>(
      args?: Subset<T, DataSetCountArgs>,
    ): Prisma.PrismaPromise<
      T extends $Utils.Record<'select', any>
        ? T['select'] extends true
          ? number
          : GetScalarType<T['select'], DataSetCountAggregateOutputType>
        : number
    >

    /**
     * Allows you to perform aggregations operations on a DataSet.
     * Note, that providing `undefined` is treated as the value not being there.
     * Read more here: https://pris.ly/d/null-undefined
     * @param {DataSetAggregateArgs} args - Select which aggregations you would like to apply and on what fields.
     * @example
     * // Ordered by age ascending
     * // Where email contains prisma.io
     * // Limited to the 10 users
     * const aggregations = await prisma.user.aggregate({
     *   _avg: {
     *     age: true,
     *   },
     *   where: {
     *     email: {
     *       contains: "prisma.io",
     *     },
     *   },
     *   orderBy: {
     *     age: "asc",
     *   },
     *   take: 10,
     * })
    **/
    aggregate<T extends DataSetAggregateArgs>(args: Subset<T, DataSetAggregateArgs>): Prisma.PrismaPromise<GetDataSetAggregateType<T>>

    /**
     * Group by DataSet.
     * Note, that providing `undefined` is treated as the value not being there.
     * Read more here: https://pris.ly/d/null-undefined
     * @param {DataSetGroupByArgs} args - Group by arguments.
     * @example
     * // Group by city, order by createdAt, get count
     * const result = await prisma.user.groupBy({
     *   by: ['city', 'createdAt'],
     *   orderBy: {
     *     createdAt: true
     *   },
     *   _count: {
     *     _all: true
     *   },
     * })
     * 
    **/
    groupBy<
      T extends DataSetGroupByArgs,
      HasSelectOrTake extends Or<
        Extends<'skip', Keys<T>>,
        Extends<'take', Keys<T>>
      >,
      OrderByArg extends True extends HasSelectOrTake
        ? { orderBy: DataSetGroupByArgs['orderBy'] }
        : { orderBy?: DataSetGroupByArgs['orderBy'] },
      OrderFields extends ExcludeUnderscoreKeys<Keys<MaybeTupleToUnion<T['orderBy']>>>,
      ByFields extends MaybeTupleToUnion<T['by']>,
      ByValid extends Has<ByFields, OrderFields>,
      HavingFields extends GetHavingFields<T['having']>,
      HavingValid extends Has<ByFields, HavingFields>,
      ByEmpty extends T['by'] extends never[] ? True : False,
      InputErrors extends ByEmpty extends True
      ? `Error: "by" must not be empty.`
      : HavingValid extends False
      ? {
          [P in HavingFields]: P extends ByFields
            ? never
            : P extends string
            ? `Error: Field "${P}" used in "having" needs to be provided in "by".`
            : [
                Error,
                'Field ',
                P,
                ` in "having" needs to be provided in "by"`,
              ]
        }[HavingFields]
      : 'take' extends Keys<T>
      ? 'orderBy' extends Keys<T>
        ? ByValid extends True
          ? {}
          : {
              [P in OrderFields]: P extends ByFields
                ? never
                : `Error: Field "${P}" in "orderBy" needs to be provided in "by"`
            }[OrderFields]
        : 'Error: If you provide "take", you also need to provide "orderBy"'
      : 'skip' extends Keys<T>
      ? 'orderBy' extends Keys<T>
        ? ByValid extends True
          ? {}
          : {
              [P in OrderFields]: P extends ByFields
                ? never
                : `Error: Field "${P}" in "orderBy" needs to be provided in "by"`
            }[OrderFields]
        : 'Error: If you provide "skip", you also need to provide "orderBy"'
      : ByValid extends True
      ? {}
      : {
          [P in OrderFields]: P extends ByFields
            ? never
            : `Error: Field "${P}" in "orderBy" needs to be provided in "by"`
        }[OrderFields]
    >(args: SubsetIntersection<T, DataSetGroupByArgs, OrderByArg> & InputErrors): {} extends InputErrors ? GetDataSetGroupByPayload<T> : Prisma.PrismaPromise<InputErrors>
  /**
   * Fields of the DataSet model
   */
  readonly fields: DataSetFieldRefs;
  }

  /**
   * The delegate class that acts as a "Promise-like" for DataSet.
   * Why is this prefixed with `Prisma__`?
   * Because we want to prevent naming conflicts as mentioned in
   * https://github.com/prisma/prisma-client-js/issues/707
   */
  export interface Prisma__DataSetClient<T, Null = never, ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> extends Prisma.PrismaPromise<T> {
    readonly [Symbol.toStringTag]: 'PrismaPromise';

    AssetPrice<T extends DataSet$AssetPriceArgs<ExtArgs> = {}>(args?: Subset<T, DataSet$AssetPriceArgs<ExtArgs>>): Prisma.PrismaPromise<$Result.GetResult<Prisma.$AssetPricePayload<ExtArgs>, T, 'findMany'> | Null>;

    /**
     * Attaches callbacks for the resolution and/or rejection of the Promise.
     * @param onfulfilled The callback to execute when the Promise is resolved.
     * @param onrejected The callback to execute when the Promise is rejected.
     * @returns A Promise for the completion of which ever callback is executed.
     */
    then<TResult1 = T, TResult2 = never>(onfulfilled?: ((value: T) => TResult1 | PromiseLike<TResult1>) | undefined | null, onrejected?: ((reason: any) => TResult2 | PromiseLike<TResult2>) | undefined | null): $Utils.JsPromise<TResult1 | TResult2>;
    /**
     * Attaches a callback for only the rejection of the Promise.
     * @param onrejected The callback to execute when the Promise is rejected.
     * @returns A Promise for the completion of the callback.
     */
    catch<TResult = never>(onrejected?: ((reason: any) => TResult | PromiseLike<TResult>) | undefined | null): $Utils.JsPromise<T | TResult>;
    /**
     * Attaches a callback that is invoked when the Promise is settled (fulfilled or rejected). The
     * resolved value cannot be modified from the callback.
     * @param onfinally The callback to execute when the Promise is settled (fulfilled or rejected).
     * @returns A Promise for the completion of the callback.
     */
    finally(onfinally?: (() => void) | undefined | null): $Utils.JsPromise<T>;
  }



  /**
   * Fields of the DataSet model
   */ 
  interface DataSetFieldRefs {
    readonly id: FieldRef<"DataSet", 'Int'>
    readonly name: FieldRef<"DataSet", 'String'>
  }
    

  // Custom InputTypes

  /**
   * DataSet findUnique
   */
  export type DataSetFindUniqueArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    /**
     * Select specific fields to fetch from the DataSet
     */
    select?: DataSetSelect<ExtArgs> | null
    /**
     * Choose, which related nodes to fetch as well.
     */
    include?: DataSetInclude<ExtArgs> | null
    /**
     * Filter, which DataSet to fetch.
     */
    where: DataSetWhereUniqueInput
  }


  /**
   * DataSet findUniqueOrThrow
   */
  export type DataSetFindUniqueOrThrowArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    /**
     * Select specific fields to fetch from the DataSet
     */
    select?: DataSetSelect<ExtArgs> | null
    /**
     * Choose, which related nodes to fetch as well.
     */
    include?: DataSetInclude<ExtArgs> | null
    /**
     * Filter, which DataSet to fetch.
     */
    where: DataSetWhereUniqueInput
  }


  /**
   * DataSet findFirst
   */
  export type DataSetFindFirstArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    /**
     * Select specific fields to fetch from the DataSet
     */
    select?: DataSetSelect<ExtArgs> | null
    /**
     * Choose, which related nodes to fetch as well.
     */
    include?: DataSetInclude<ExtArgs> | null
    /**
     * Filter, which DataSet to fetch.
     */
    where?: DataSetWhereInput
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/sorting Sorting Docs}
     * 
     * Determine the order of DataSets to fetch.
     */
    orderBy?: DataSetOrderByWithRelationInput | DataSetOrderByWithRelationInput[]
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/pagination#cursor-based-pagination Cursor Docs}
     * 
     * Sets the position for searching for DataSets.
     */
    cursor?: DataSetWhereUniqueInput
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/pagination Pagination Docs}
     * 
     * Take `±n` DataSets from the position of the cursor.
     */
    take?: number
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/pagination Pagination Docs}
     * 
     * Skip the first `n` DataSets.
     */
    skip?: number
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/distinct Distinct Docs}
     * 
     * Filter by unique combinations of DataSets.
     */
    distinct?: DataSetScalarFieldEnum | DataSetScalarFieldEnum[]
  }


  /**
   * DataSet findFirstOrThrow
   */
  export type DataSetFindFirstOrThrowArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    /**
     * Select specific fields to fetch from the DataSet
     */
    select?: DataSetSelect<ExtArgs> | null
    /**
     * Choose, which related nodes to fetch as well.
     */
    include?: DataSetInclude<ExtArgs> | null
    /**
     * Filter, which DataSet to fetch.
     */
    where?: DataSetWhereInput
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/sorting Sorting Docs}
     * 
     * Determine the order of DataSets to fetch.
     */
    orderBy?: DataSetOrderByWithRelationInput | DataSetOrderByWithRelationInput[]
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/pagination#cursor-based-pagination Cursor Docs}
     * 
     * Sets the position for searching for DataSets.
     */
    cursor?: DataSetWhereUniqueInput
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/pagination Pagination Docs}
     * 
     * Take `±n` DataSets from the position of the cursor.
     */
    take?: number
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/pagination Pagination Docs}
     * 
     * Skip the first `n` DataSets.
     */
    skip?: number
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/distinct Distinct Docs}
     * 
     * Filter by unique combinations of DataSets.
     */
    distinct?: DataSetScalarFieldEnum | DataSetScalarFieldEnum[]
  }


  /**
   * DataSet findMany
   */
  export type DataSetFindManyArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    /**
     * Select specific fields to fetch from the DataSet
     */
    select?: DataSetSelect<ExtArgs> | null
    /**
     * Choose, which related nodes to fetch as well.
     */
    include?: DataSetInclude<ExtArgs> | null
    /**
     * Filter, which DataSets to fetch.
     */
    where?: DataSetWhereInput
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/sorting Sorting Docs}
     * 
     * Determine the order of DataSets to fetch.
     */
    orderBy?: DataSetOrderByWithRelationInput | DataSetOrderByWithRelationInput[]
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/pagination#cursor-based-pagination Cursor Docs}
     * 
     * Sets the position for listing DataSets.
     */
    cursor?: DataSetWhereUniqueInput
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/pagination Pagination Docs}
     * 
     * Take `±n` DataSets from the position of the cursor.
     */
    take?: number
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/pagination Pagination Docs}
     * 
     * Skip the first `n` DataSets.
     */
    skip?: number
    distinct?: DataSetScalarFieldEnum | DataSetScalarFieldEnum[]
  }


  /**
   * DataSet create
   */
  export type DataSetCreateArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    /**
     * Select specific fields to fetch from the DataSet
     */
    select?: DataSetSelect<ExtArgs> | null
    /**
     * Choose, which related nodes to fetch as well.
     */
    include?: DataSetInclude<ExtArgs> | null
    /**
     * The data needed to create a DataSet.
     */
    data: XOR<DataSetCreateInput, DataSetUncheckedCreateInput>
  }


  /**
   * DataSet createMany
   */
  export type DataSetCreateManyArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    /**
     * The data used to create many DataSets.
     */
    data: DataSetCreateManyInput | DataSetCreateManyInput[]
    skipDuplicates?: boolean
  }


  /**
   * DataSet update
   */
  export type DataSetUpdateArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    /**
     * Select specific fields to fetch from the DataSet
     */
    select?: DataSetSelect<ExtArgs> | null
    /**
     * Choose, which related nodes to fetch as well.
     */
    include?: DataSetInclude<ExtArgs> | null
    /**
     * The data needed to update a DataSet.
     */
    data: XOR<DataSetUpdateInput, DataSetUncheckedUpdateInput>
    /**
     * Choose, which DataSet to update.
     */
    where: DataSetWhereUniqueInput
  }


  /**
   * DataSet updateMany
   */
  export type DataSetUpdateManyArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    /**
     * The data used to update DataSets.
     */
    data: XOR<DataSetUpdateManyMutationInput, DataSetUncheckedUpdateManyInput>
    /**
     * Filter which DataSets to update
     */
    where?: DataSetWhereInput
  }


  /**
   * DataSet upsert
   */
  export type DataSetUpsertArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    /**
     * Select specific fields to fetch from the DataSet
     */
    select?: DataSetSelect<ExtArgs> | null
    /**
     * Choose, which related nodes to fetch as well.
     */
    include?: DataSetInclude<ExtArgs> | null
    /**
     * The filter to search for the DataSet to update in case it exists.
     */
    where: DataSetWhereUniqueInput
    /**
     * In case the DataSet found by the `where` argument doesn't exist, create a new DataSet with this data.
     */
    create: XOR<DataSetCreateInput, DataSetUncheckedCreateInput>
    /**
     * In case the DataSet was found with the provided `where` argument, update it with this data.
     */
    update: XOR<DataSetUpdateInput, DataSetUncheckedUpdateInput>
  }


  /**
   * DataSet delete
   */
  export type DataSetDeleteArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    /**
     * Select specific fields to fetch from the DataSet
     */
    select?: DataSetSelect<ExtArgs> | null
    /**
     * Choose, which related nodes to fetch as well.
     */
    include?: DataSetInclude<ExtArgs> | null
    /**
     * Filter which DataSet to delete.
     */
    where: DataSetWhereUniqueInput
  }


  /**
   * DataSet deleteMany
   */
  export type DataSetDeleteManyArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    /**
     * Filter which DataSets to delete
     */
    where?: DataSetWhereInput
  }


  /**
   * DataSet.AssetPrice
   */
  export type DataSet$AssetPriceArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    /**
     * Select specific fields to fetch from the AssetPrice
     */
    select?: AssetPriceSelect<ExtArgs> | null
    /**
     * Choose, which related nodes to fetch as well.
     */
    include?: AssetPriceInclude<ExtArgs> | null
    where?: AssetPriceWhereInput
    orderBy?: AssetPriceOrderByWithRelationInput | AssetPriceOrderByWithRelationInput[]
    cursor?: AssetPriceWhereUniqueInput
    take?: number
    skip?: number
    distinct?: AssetPriceScalarFieldEnum | AssetPriceScalarFieldEnum[]
  }


  /**
   * DataSet without action
   */
  export type DataSetDefaultArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    /**
     * Select specific fields to fetch from the DataSet
     */
    select?: DataSetSelect<ExtArgs> | null
    /**
     * Choose, which related nodes to fetch as well.
     */
    include?: DataSetInclude<ExtArgs> | null
  }



  /**
   * Model Signer
   */

  export type AggregateSigner = {
    _count: SignerCountAggregateOutputType | null
    _avg: SignerAvgAggregateOutputType | null
    _sum: SignerSumAggregateOutputType | null
    _min: SignerMinAggregateOutputType | null
    _max: SignerMaxAggregateOutputType | null
  }

  export type SignerAvgAggregateOutputType = {
    id: number | null
  }

  export type SignerSumAggregateOutputType = {
    id: number | null
  }

  export type SignerMinAggregateOutputType = {
    id: number | null
    key: string | null
    name: string | null
  }

  export type SignerMaxAggregateOutputType = {
    id: number | null
    key: string | null
    name: string | null
  }

  export type SignerCountAggregateOutputType = {
    id: number
    key: number
    name: number
    _all: number
  }


  export type SignerAvgAggregateInputType = {
    id?: true
  }

  export type SignerSumAggregateInputType = {
    id?: true
  }

  export type SignerMinAggregateInputType = {
    id?: true
    key?: true
    name?: true
  }

  export type SignerMaxAggregateInputType = {
    id?: true
    key?: true
    name?: true
  }

  export type SignerCountAggregateInputType = {
    id?: true
    key?: true
    name?: true
    _all?: true
  }

  export type SignerAggregateArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    /**
     * Filter which Signer to aggregate.
     */
    where?: SignerWhereInput
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/sorting Sorting Docs}
     * 
     * Determine the order of Signers to fetch.
     */
    orderBy?: SignerOrderByWithRelationInput | SignerOrderByWithRelationInput[]
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/pagination#cursor-based-pagination Cursor Docs}
     * 
     * Sets the start position
     */
    cursor?: SignerWhereUniqueInput
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/pagination Pagination Docs}
     * 
     * Take `±n` Signers from the position of the cursor.
     */
    take?: number
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/pagination Pagination Docs}
     * 
     * Skip the first `n` Signers.
     */
    skip?: number
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/aggregations Aggregation Docs}
     * 
     * Count returned Signers
    **/
    _count?: true | SignerCountAggregateInputType
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/aggregations Aggregation Docs}
     * 
     * Select which fields to average
    **/
    _avg?: SignerAvgAggregateInputType
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/aggregations Aggregation Docs}
     * 
     * Select which fields to sum
    **/
    _sum?: SignerSumAggregateInputType
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/aggregations Aggregation Docs}
     * 
     * Select which fields to find the minimum value
    **/
    _min?: SignerMinAggregateInputType
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/aggregations Aggregation Docs}
     * 
     * Select which fields to find the maximum value
    **/
    _max?: SignerMaxAggregateInputType
  }

  export type GetSignerAggregateType<T extends SignerAggregateArgs> = {
        [P in keyof T & keyof AggregateSigner]: P extends '_count' | 'count'
      ? T[P] extends true
        ? number
        : GetScalarType<T[P], AggregateSigner[P]>
      : GetScalarType<T[P], AggregateSigner[P]>
  }




  export type SignerGroupByArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    where?: SignerWhereInput
    orderBy?: SignerOrderByWithAggregationInput | SignerOrderByWithAggregationInput[]
    by: SignerScalarFieldEnum[] | SignerScalarFieldEnum
    having?: SignerScalarWhereWithAggregatesInput
    take?: number
    skip?: number
    _count?: SignerCountAggregateInputType | true
    _avg?: SignerAvgAggregateInputType
    _sum?: SignerSumAggregateInputType
    _min?: SignerMinAggregateInputType
    _max?: SignerMaxAggregateInputType
  }

  export type SignerGroupByOutputType = {
    id: number
    key: string
    name: string | null
    _count: SignerCountAggregateOutputType | null
    _avg: SignerAvgAggregateOutputType | null
    _sum: SignerSumAggregateOutputType | null
    _min: SignerMinAggregateOutputType | null
    _max: SignerMaxAggregateOutputType | null
  }

  type GetSignerGroupByPayload<T extends SignerGroupByArgs> = Prisma.PrismaPromise<
    Array<
      PickEnumerable<SignerGroupByOutputType, T['by']> &
        {
          [P in ((keyof T) & (keyof SignerGroupByOutputType))]: P extends '_count'
            ? T[P] extends boolean
              ? number
              : GetScalarType<T[P], SignerGroupByOutputType[P]>
            : GetScalarType<T[P], SignerGroupByOutputType[P]>
        }
      >
    >


  export type SignerSelect<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = $Extensions.GetSelect<{
    id?: boolean
    key?: boolean
    name?: boolean
    signersOnAssetPrice?: boolean | Signer$signersOnAssetPriceArgs<ExtArgs>
    _count?: boolean | SignerCountOutputTypeDefaultArgs<ExtArgs>
  }, ExtArgs["result"]["signer"]>

  export type SignerSelectScalar = {
    id?: boolean
    key?: boolean
    name?: boolean
  }

  export type SignerInclude<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    signersOnAssetPrice?: boolean | Signer$signersOnAssetPriceArgs<ExtArgs>
    _count?: boolean | SignerCountOutputTypeDefaultArgs<ExtArgs>
  }


  export type $SignerPayload<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    name: "Signer"
    objects: {
      signersOnAssetPrice: Prisma.$SignersOnAssetPricePayload<ExtArgs>[]
    }
    scalars: $Extensions.GetPayloadResult<{
      id: number
      key: string
      name: string | null
    }, ExtArgs["result"]["signer"]>
    composites: {}
  }


  type SignerGetPayload<S extends boolean | null | undefined | SignerDefaultArgs> = $Result.GetResult<Prisma.$SignerPayload, S>

  type SignerCountArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = 
    Omit<SignerFindManyArgs, 'select' | 'include' | 'distinct' > & {
      select?: SignerCountAggregateInputType | true
    }

  export interface SignerDelegate<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> {
    [K: symbol]: { types: Prisma.TypeMap<ExtArgs>['model']['Signer'], meta: { name: 'Signer' } }
    /**
     * Find zero or one Signer that matches the filter.
     * @param {SignerFindUniqueArgs} args - Arguments to find a Signer
     * @example
     * // Get one Signer
     * const signer = await prisma.signer.findUnique({
     *   where: {
     *     // ... provide filter here
     *   }
     * })
    **/
    findUnique<T extends SignerFindUniqueArgs<ExtArgs>>(
      args: SelectSubset<T, SignerFindUniqueArgs<ExtArgs>>
    ): Prisma__SignerClient<$Result.GetResult<Prisma.$SignerPayload<ExtArgs>, T, 'findUnique'> | null, null, ExtArgs>

    /**
     * Find one Signer that matches the filter or throw an error  with `error.code='P2025'` 
     *     if no matches were found.
     * @param {SignerFindUniqueOrThrowArgs} args - Arguments to find a Signer
     * @example
     * // Get one Signer
     * const signer = await prisma.signer.findUniqueOrThrow({
     *   where: {
     *     // ... provide filter here
     *   }
     * })
    **/
    findUniqueOrThrow<T extends SignerFindUniqueOrThrowArgs<ExtArgs>>(
      args?: SelectSubset<T, SignerFindUniqueOrThrowArgs<ExtArgs>>
    ): Prisma__SignerClient<$Result.GetResult<Prisma.$SignerPayload<ExtArgs>, T, 'findUniqueOrThrow'>, never, ExtArgs>

    /**
     * Find the first Signer that matches the filter.
     * Note, that providing `undefined` is treated as the value not being there.
     * Read more here: https://pris.ly/d/null-undefined
     * @param {SignerFindFirstArgs} args - Arguments to find a Signer
     * @example
     * // Get one Signer
     * const signer = await prisma.signer.findFirst({
     *   where: {
     *     // ... provide filter here
     *   }
     * })
    **/
    findFirst<T extends SignerFindFirstArgs<ExtArgs>>(
      args?: SelectSubset<T, SignerFindFirstArgs<ExtArgs>>
    ): Prisma__SignerClient<$Result.GetResult<Prisma.$SignerPayload<ExtArgs>, T, 'findFirst'> | null, null, ExtArgs>

    /**
     * Find the first Signer that matches the filter or
     * throw `PrismaKnownClientError` with `P2025` code if no matches were found.
     * Note, that providing `undefined` is treated as the value not being there.
     * Read more here: https://pris.ly/d/null-undefined
     * @param {SignerFindFirstOrThrowArgs} args - Arguments to find a Signer
     * @example
     * // Get one Signer
     * const signer = await prisma.signer.findFirstOrThrow({
     *   where: {
     *     // ... provide filter here
     *   }
     * })
    **/
    findFirstOrThrow<T extends SignerFindFirstOrThrowArgs<ExtArgs>>(
      args?: SelectSubset<T, SignerFindFirstOrThrowArgs<ExtArgs>>
    ): Prisma__SignerClient<$Result.GetResult<Prisma.$SignerPayload<ExtArgs>, T, 'findFirstOrThrow'>, never, ExtArgs>

    /**
     * Find zero or more Signers that matches the filter.
     * Note, that providing `undefined` is treated as the value not being there.
     * Read more here: https://pris.ly/d/null-undefined
     * @param {SignerFindManyArgs=} args - Arguments to filter and select certain fields only.
     * @example
     * // Get all Signers
     * const signers = await prisma.signer.findMany()
     * 
     * // Get first 10 Signers
     * const signers = await prisma.signer.findMany({ take: 10 })
     * 
     * // Only select the `id`
     * const signerWithIdOnly = await prisma.signer.findMany({ select: { id: true } })
     * 
    **/
    findMany<T extends SignerFindManyArgs<ExtArgs>>(
      args?: SelectSubset<T, SignerFindManyArgs<ExtArgs>>
    ): Prisma.PrismaPromise<$Result.GetResult<Prisma.$SignerPayload<ExtArgs>, T, 'findMany'>>

    /**
     * Create a Signer.
     * @param {SignerCreateArgs} args - Arguments to create a Signer.
     * @example
     * // Create one Signer
     * const Signer = await prisma.signer.create({
     *   data: {
     *     // ... data to create a Signer
     *   }
     * })
     * 
    **/
    create<T extends SignerCreateArgs<ExtArgs>>(
      args: SelectSubset<T, SignerCreateArgs<ExtArgs>>
    ): Prisma__SignerClient<$Result.GetResult<Prisma.$SignerPayload<ExtArgs>, T, 'create'>, never, ExtArgs>

    /**
     * Create many Signers.
     *     @param {SignerCreateManyArgs} args - Arguments to create many Signers.
     *     @example
     *     // Create many Signers
     *     const signer = await prisma.signer.createMany({
     *       data: {
     *         // ... provide data here
     *       }
     *     })
     *     
    **/
    createMany<T extends SignerCreateManyArgs<ExtArgs>>(
      args?: SelectSubset<T, SignerCreateManyArgs<ExtArgs>>
    ): Prisma.PrismaPromise<BatchPayload>

    /**
     * Delete a Signer.
     * @param {SignerDeleteArgs} args - Arguments to delete one Signer.
     * @example
     * // Delete one Signer
     * const Signer = await prisma.signer.delete({
     *   where: {
     *     // ... filter to delete one Signer
     *   }
     * })
     * 
    **/
    delete<T extends SignerDeleteArgs<ExtArgs>>(
      args: SelectSubset<T, SignerDeleteArgs<ExtArgs>>
    ): Prisma__SignerClient<$Result.GetResult<Prisma.$SignerPayload<ExtArgs>, T, 'delete'>, never, ExtArgs>

    /**
     * Update one Signer.
     * @param {SignerUpdateArgs} args - Arguments to update one Signer.
     * @example
     * // Update one Signer
     * const signer = await prisma.signer.update({
     *   where: {
     *     // ... provide filter here
     *   },
     *   data: {
     *     // ... provide data here
     *   }
     * })
     * 
    **/
    update<T extends SignerUpdateArgs<ExtArgs>>(
      args: SelectSubset<T, SignerUpdateArgs<ExtArgs>>
    ): Prisma__SignerClient<$Result.GetResult<Prisma.$SignerPayload<ExtArgs>, T, 'update'>, never, ExtArgs>

    /**
     * Delete zero or more Signers.
     * @param {SignerDeleteManyArgs} args - Arguments to filter Signers to delete.
     * @example
     * // Delete a few Signers
     * const { count } = await prisma.signer.deleteMany({
     *   where: {
     *     // ... provide filter here
     *   }
     * })
     * 
    **/
    deleteMany<T extends SignerDeleteManyArgs<ExtArgs>>(
      args?: SelectSubset<T, SignerDeleteManyArgs<ExtArgs>>
    ): Prisma.PrismaPromise<BatchPayload>

    /**
     * Update zero or more Signers.
     * Note, that providing `undefined` is treated as the value not being there.
     * Read more here: https://pris.ly/d/null-undefined
     * @param {SignerUpdateManyArgs} args - Arguments to update one or more rows.
     * @example
     * // Update many Signers
     * const signer = await prisma.signer.updateMany({
     *   where: {
     *     // ... provide filter here
     *   },
     *   data: {
     *     // ... provide data here
     *   }
     * })
     * 
    **/
    updateMany<T extends SignerUpdateManyArgs<ExtArgs>>(
      args: SelectSubset<T, SignerUpdateManyArgs<ExtArgs>>
    ): Prisma.PrismaPromise<BatchPayload>

    /**
     * Create or update one Signer.
     * @param {SignerUpsertArgs} args - Arguments to update or create a Signer.
     * @example
     * // Update or create a Signer
     * const signer = await prisma.signer.upsert({
     *   create: {
     *     // ... data to create a Signer
     *   },
     *   update: {
     *     // ... in case it already exists, update
     *   },
     *   where: {
     *     // ... the filter for the Signer we want to update
     *   }
     * })
    **/
    upsert<T extends SignerUpsertArgs<ExtArgs>>(
      args: SelectSubset<T, SignerUpsertArgs<ExtArgs>>
    ): Prisma__SignerClient<$Result.GetResult<Prisma.$SignerPayload<ExtArgs>, T, 'upsert'>, never, ExtArgs>

    /**
     * Count the number of Signers.
     * Note, that providing `undefined` is treated as the value not being there.
     * Read more here: https://pris.ly/d/null-undefined
     * @param {SignerCountArgs} args - Arguments to filter Signers to count.
     * @example
     * // Count the number of Signers
     * const count = await prisma.signer.count({
     *   where: {
     *     // ... the filter for the Signers we want to count
     *   }
     * })
    **/
    count<T extends SignerCountArgs>(
      args?: Subset<T, SignerCountArgs>,
    ): Prisma.PrismaPromise<
      T extends $Utils.Record<'select', any>
        ? T['select'] extends true
          ? number
          : GetScalarType<T['select'], SignerCountAggregateOutputType>
        : number
    >

    /**
     * Allows you to perform aggregations operations on a Signer.
     * Note, that providing `undefined` is treated as the value not being there.
     * Read more here: https://pris.ly/d/null-undefined
     * @param {SignerAggregateArgs} args - Select which aggregations you would like to apply and on what fields.
     * @example
     * // Ordered by age ascending
     * // Where email contains prisma.io
     * // Limited to the 10 users
     * const aggregations = await prisma.user.aggregate({
     *   _avg: {
     *     age: true,
     *   },
     *   where: {
     *     email: {
     *       contains: "prisma.io",
     *     },
     *   },
     *   orderBy: {
     *     age: "asc",
     *   },
     *   take: 10,
     * })
    **/
    aggregate<T extends SignerAggregateArgs>(args: Subset<T, SignerAggregateArgs>): Prisma.PrismaPromise<GetSignerAggregateType<T>>

    /**
     * Group by Signer.
     * Note, that providing `undefined` is treated as the value not being there.
     * Read more here: https://pris.ly/d/null-undefined
     * @param {SignerGroupByArgs} args - Group by arguments.
     * @example
     * // Group by city, order by createdAt, get count
     * const result = await prisma.user.groupBy({
     *   by: ['city', 'createdAt'],
     *   orderBy: {
     *     createdAt: true
     *   },
     *   _count: {
     *     _all: true
     *   },
     * })
     * 
    **/
    groupBy<
      T extends SignerGroupByArgs,
      HasSelectOrTake extends Or<
        Extends<'skip', Keys<T>>,
        Extends<'take', Keys<T>>
      >,
      OrderByArg extends True extends HasSelectOrTake
        ? { orderBy: SignerGroupByArgs['orderBy'] }
        : { orderBy?: SignerGroupByArgs['orderBy'] },
      OrderFields extends ExcludeUnderscoreKeys<Keys<MaybeTupleToUnion<T['orderBy']>>>,
      ByFields extends MaybeTupleToUnion<T['by']>,
      ByValid extends Has<ByFields, OrderFields>,
      HavingFields extends GetHavingFields<T['having']>,
      HavingValid extends Has<ByFields, HavingFields>,
      ByEmpty extends T['by'] extends never[] ? True : False,
      InputErrors extends ByEmpty extends True
      ? `Error: "by" must not be empty.`
      : HavingValid extends False
      ? {
          [P in HavingFields]: P extends ByFields
            ? never
            : P extends string
            ? `Error: Field "${P}" used in "having" needs to be provided in "by".`
            : [
                Error,
                'Field ',
                P,
                ` in "having" needs to be provided in "by"`,
              ]
        }[HavingFields]
      : 'take' extends Keys<T>
      ? 'orderBy' extends Keys<T>
        ? ByValid extends True
          ? {}
          : {
              [P in OrderFields]: P extends ByFields
                ? never
                : `Error: Field "${P}" in "orderBy" needs to be provided in "by"`
            }[OrderFields]
        : 'Error: If you provide "take", you also need to provide "orderBy"'
      : 'skip' extends Keys<T>
      ? 'orderBy' extends Keys<T>
        ? ByValid extends True
          ? {}
          : {
              [P in OrderFields]: P extends ByFields
                ? never
                : `Error: Field "${P}" in "orderBy" needs to be provided in "by"`
            }[OrderFields]
        : 'Error: If you provide "skip", you also need to provide "orderBy"'
      : ByValid extends True
      ? {}
      : {
          [P in OrderFields]: P extends ByFields
            ? never
            : `Error: Field "${P}" in "orderBy" needs to be provided in "by"`
        }[OrderFields]
    >(args: SubsetIntersection<T, SignerGroupByArgs, OrderByArg> & InputErrors): {} extends InputErrors ? GetSignerGroupByPayload<T> : Prisma.PrismaPromise<InputErrors>
  /**
   * Fields of the Signer model
   */
  readonly fields: SignerFieldRefs;
  }

  /**
   * The delegate class that acts as a "Promise-like" for Signer.
   * Why is this prefixed with `Prisma__`?
   * Because we want to prevent naming conflicts as mentioned in
   * https://github.com/prisma/prisma-client-js/issues/707
   */
  export interface Prisma__SignerClient<T, Null = never, ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> extends Prisma.PrismaPromise<T> {
    readonly [Symbol.toStringTag]: 'PrismaPromise';

    signersOnAssetPrice<T extends Signer$signersOnAssetPriceArgs<ExtArgs> = {}>(args?: Subset<T, Signer$signersOnAssetPriceArgs<ExtArgs>>): Prisma.PrismaPromise<$Result.GetResult<Prisma.$SignersOnAssetPricePayload<ExtArgs>, T, 'findMany'> | Null>;

    /**
     * Attaches callbacks for the resolution and/or rejection of the Promise.
     * @param onfulfilled The callback to execute when the Promise is resolved.
     * @param onrejected The callback to execute when the Promise is rejected.
     * @returns A Promise for the completion of which ever callback is executed.
     */
    then<TResult1 = T, TResult2 = never>(onfulfilled?: ((value: T) => TResult1 | PromiseLike<TResult1>) | undefined | null, onrejected?: ((reason: any) => TResult2 | PromiseLike<TResult2>) | undefined | null): $Utils.JsPromise<TResult1 | TResult2>;
    /**
     * Attaches a callback for only the rejection of the Promise.
     * @param onrejected The callback to execute when the Promise is rejected.
     * @returns A Promise for the completion of the callback.
     */
    catch<TResult = never>(onrejected?: ((reason: any) => TResult | PromiseLike<TResult>) | undefined | null): $Utils.JsPromise<T | TResult>;
    /**
     * Attaches a callback that is invoked when the Promise is settled (fulfilled or rejected). The
     * resolved value cannot be modified from the callback.
     * @param onfinally The callback to execute when the Promise is settled (fulfilled or rejected).
     * @returns A Promise for the completion of the callback.
     */
    finally(onfinally?: (() => void) | undefined | null): $Utils.JsPromise<T>;
  }



  /**
   * Fields of the Signer model
   */ 
  interface SignerFieldRefs {
    readonly id: FieldRef<"Signer", 'Int'>
    readonly key: FieldRef<"Signer", 'String'>
    readonly name: FieldRef<"Signer", 'String'>
  }
    

  // Custom InputTypes

  /**
   * Signer findUnique
   */
  export type SignerFindUniqueArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    /**
     * Select specific fields to fetch from the Signer
     */
    select?: SignerSelect<ExtArgs> | null
    /**
     * Choose, which related nodes to fetch as well.
     */
    include?: SignerInclude<ExtArgs> | null
    /**
     * Filter, which Signer to fetch.
     */
    where: SignerWhereUniqueInput
  }


  /**
   * Signer findUniqueOrThrow
   */
  export type SignerFindUniqueOrThrowArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    /**
     * Select specific fields to fetch from the Signer
     */
    select?: SignerSelect<ExtArgs> | null
    /**
     * Choose, which related nodes to fetch as well.
     */
    include?: SignerInclude<ExtArgs> | null
    /**
     * Filter, which Signer to fetch.
     */
    where: SignerWhereUniqueInput
  }


  /**
   * Signer findFirst
   */
  export type SignerFindFirstArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    /**
     * Select specific fields to fetch from the Signer
     */
    select?: SignerSelect<ExtArgs> | null
    /**
     * Choose, which related nodes to fetch as well.
     */
    include?: SignerInclude<ExtArgs> | null
    /**
     * Filter, which Signer to fetch.
     */
    where?: SignerWhereInput
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/sorting Sorting Docs}
     * 
     * Determine the order of Signers to fetch.
     */
    orderBy?: SignerOrderByWithRelationInput | SignerOrderByWithRelationInput[]
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/pagination#cursor-based-pagination Cursor Docs}
     * 
     * Sets the position for searching for Signers.
     */
    cursor?: SignerWhereUniqueInput
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/pagination Pagination Docs}
     * 
     * Take `±n` Signers from the position of the cursor.
     */
    take?: number
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/pagination Pagination Docs}
     * 
     * Skip the first `n` Signers.
     */
    skip?: number
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/distinct Distinct Docs}
     * 
     * Filter by unique combinations of Signers.
     */
    distinct?: SignerScalarFieldEnum | SignerScalarFieldEnum[]
  }


  /**
   * Signer findFirstOrThrow
   */
  export type SignerFindFirstOrThrowArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    /**
     * Select specific fields to fetch from the Signer
     */
    select?: SignerSelect<ExtArgs> | null
    /**
     * Choose, which related nodes to fetch as well.
     */
    include?: SignerInclude<ExtArgs> | null
    /**
     * Filter, which Signer to fetch.
     */
    where?: SignerWhereInput
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/sorting Sorting Docs}
     * 
     * Determine the order of Signers to fetch.
     */
    orderBy?: SignerOrderByWithRelationInput | SignerOrderByWithRelationInput[]
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/pagination#cursor-based-pagination Cursor Docs}
     * 
     * Sets the position for searching for Signers.
     */
    cursor?: SignerWhereUniqueInput
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/pagination Pagination Docs}
     * 
     * Take `±n` Signers from the position of the cursor.
     */
    take?: number
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/pagination Pagination Docs}
     * 
     * Skip the first `n` Signers.
     */
    skip?: number
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/distinct Distinct Docs}
     * 
     * Filter by unique combinations of Signers.
     */
    distinct?: SignerScalarFieldEnum | SignerScalarFieldEnum[]
  }


  /**
   * Signer findMany
   */
  export type SignerFindManyArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    /**
     * Select specific fields to fetch from the Signer
     */
    select?: SignerSelect<ExtArgs> | null
    /**
     * Choose, which related nodes to fetch as well.
     */
    include?: SignerInclude<ExtArgs> | null
    /**
     * Filter, which Signers to fetch.
     */
    where?: SignerWhereInput
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/sorting Sorting Docs}
     * 
     * Determine the order of Signers to fetch.
     */
    orderBy?: SignerOrderByWithRelationInput | SignerOrderByWithRelationInput[]
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/pagination#cursor-based-pagination Cursor Docs}
     * 
     * Sets the position for listing Signers.
     */
    cursor?: SignerWhereUniqueInput
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/pagination Pagination Docs}
     * 
     * Take `±n` Signers from the position of the cursor.
     */
    take?: number
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/pagination Pagination Docs}
     * 
     * Skip the first `n` Signers.
     */
    skip?: number
    distinct?: SignerScalarFieldEnum | SignerScalarFieldEnum[]
  }


  /**
   * Signer create
   */
  export type SignerCreateArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    /**
     * Select specific fields to fetch from the Signer
     */
    select?: SignerSelect<ExtArgs> | null
    /**
     * Choose, which related nodes to fetch as well.
     */
    include?: SignerInclude<ExtArgs> | null
    /**
     * The data needed to create a Signer.
     */
    data: XOR<SignerCreateInput, SignerUncheckedCreateInput>
  }


  /**
   * Signer createMany
   */
  export type SignerCreateManyArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    /**
     * The data used to create many Signers.
     */
    data: SignerCreateManyInput | SignerCreateManyInput[]
    skipDuplicates?: boolean
  }


  /**
   * Signer update
   */
  export type SignerUpdateArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    /**
     * Select specific fields to fetch from the Signer
     */
    select?: SignerSelect<ExtArgs> | null
    /**
     * Choose, which related nodes to fetch as well.
     */
    include?: SignerInclude<ExtArgs> | null
    /**
     * The data needed to update a Signer.
     */
    data: XOR<SignerUpdateInput, SignerUncheckedUpdateInput>
    /**
     * Choose, which Signer to update.
     */
    where: SignerWhereUniqueInput
  }


  /**
   * Signer updateMany
   */
  export type SignerUpdateManyArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    /**
     * The data used to update Signers.
     */
    data: XOR<SignerUpdateManyMutationInput, SignerUncheckedUpdateManyInput>
    /**
     * Filter which Signers to update
     */
    where?: SignerWhereInput
  }


  /**
   * Signer upsert
   */
  export type SignerUpsertArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    /**
     * Select specific fields to fetch from the Signer
     */
    select?: SignerSelect<ExtArgs> | null
    /**
     * Choose, which related nodes to fetch as well.
     */
    include?: SignerInclude<ExtArgs> | null
    /**
     * The filter to search for the Signer to update in case it exists.
     */
    where: SignerWhereUniqueInput
    /**
     * In case the Signer found by the `where` argument doesn't exist, create a new Signer with this data.
     */
    create: XOR<SignerCreateInput, SignerUncheckedCreateInput>
    /**
     * In case the Signer was found with the provided `where` argument, update it with this data.
     */
    update: XOR<SignerUpdateInput, SignerUncheckedUpdateInput>
  }


  /**
   * Signer delete
   */
  export type SignerDeleteArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    /**
     * Select specific fields to fetch from the Signer
     */
    select?: SignerSelect<ExtArgs> | null
    /**
     * Choose, which related nodes to fetch as well.
     */
    include?: SignerInclude<ExtArgs> | null
    /**
     * Filter which Signer to delete.
     */
    where: SignerWhereUniqueInput
  }


  /**
   * Signer deleteMany
   */
  export type SignerDeleteManyArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    /**
     * Filter which Signers to delete
     */
    where?: SignerWhereInput
  }


  /**
   * Signer.signersOnAssetPrice
   */
  export type Signer$signersOnAssetPriceArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    /**
     * Select specific fields to fetch from the SignersOnAssetPrice
     */
    select?: SignersOnAssetPriceSelect<ExtArgs> | null
    /**
     * Choose, which related nodes to fetch as well.
     */
    include?: SignersOnAssetPriceInclude<ExtArgs> | null
    where?: SignersOnAssetPriceWhereInput
    orderBy?: SignersOnAssetPriceOrderByWithRelationInput | SignersOnAssetPriceOrderByWithRelationInput[]
    cursor?: SignersOnAssetPriceWhereUniqueInput
    take?: number
    skip?: number
    distinct?: SignersOnAssetPriceScalarFieldEnum | SignersOnAssetPriceScalarFieldEnum[]
  }


  /**
   * Signer without action
   */
  export type SignerDefaultArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    /**
     * Select specific fields to fetch from the Signer
     */
    select?: SignerSelect<ExtArgs> | null
    /**
     * Choose, which related nodes to fetch as well.
     */
    include?: SignerInclude<ExtArgs> | null
  }



  /**
   * Model SignersOnAssetPrice
   */

  export type AggregateSignersOnAssetPrice = {
    _count: SignersOnAssetPriceCountAggregateOutputType | null
    _avg: SignersOnAssetPriceAvgAggregateOutputType | null
    _sum: SignersOnAssetPriceSumAggregateOutputType | null
    _min: SignersOnAssetPriceMinAggregateOutputType | null
    _max: SignersOnAssetPriceMaxAggregateOutputType | null
  }

  export type SignersOnAssetPriceAvgAggregateOutputType = {
    signerId: number | null
    assetPriceId: number | null
  }

  export type SignersOnAssetPriceSumAggregateOutputType = {
    signerId: number | null
    assetPriceId: number | null
  }

  export type SignersOnAssetPriceMinAggregateOutputType = {
    signerId: number | null
    assetPriceId: number | null
  }

  export type SignersOnAssetPriceMaxAggregateOutputType = {
    signerId: number | null
    assetPriceId: number | null
  }

  export type SignersOnAssetPriceCountAggregateOutputType = {
    signerId: number
    assetPriceId: number
    _all: number
  }


  export type SignersOnAssetPriceAvgAggregateInputType = {
    signerId?: true
    assetPriceId?: true
  }

  export type SignersOnAssetPriceSumAggregateInputType = {
    signerId?: true
    assetPriceId?: true
  }

  export type SignersOnAssetPriceMinAggregateInputType = {
    signerId?: true
    assetPriceId?: true
  }

  export type SignersOnAssetPriceMaxAggregateInputType = {
    signerId?: true
    assetPriceId?: true
  }

  export type SignersOnAssetPriceCountAggregateInputType = {
    signerId?: true
    assetPriceId?: true
    _all?: true
  }

  export type SignersOnAssetPriceAggregateArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    /**
     * Filter which SignersOnAssetPrice to aggregate.
     */
    where?: SignersOnAssetPriceWhereInput
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/sorting Sorting Docs}
     * 
     * Determine the order of SignersOnAssetPrices to fetch.
     */
    orderBy?: SignersOnAssetPriceOrderByWithRelationInput | SignersOnAssetPriceOrderByWithRelationInput[]
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/pagination#cursor-based-pagination Cursor Docs}
     * 
     * Sets the start position
     */
    cursor?: SignersOnAssetPriceWhereUniqueInput
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/pagination Pagination Docs}
     * 
     * Take `±n` SignersOnAssetPrices from the position of the cursor.
     */
    take?: number
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/pagination Pagination Docs}
     * 
     * Skip the first `n` SignersOnAssetPrices.
     */
    skip?: number
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/aggregations Aggregation Docs}
     * 
     * Count returned SignersOnAssetPrices
    **/
    _count?: true | SignersOnAssetPriceCountAggregateInputType
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/aggregations Aggregation Docs}
     * 
     * Select which fields to average
    **/
    _avg?: SignersOnAssetPriceAvgAggregateInputType
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/aggregations Aggregation Docs}
     * 
     * Select which fields to sum
    **/
    _sum?: SignersOnAssetPriceSumAggregateInputType
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/aggregations Aggregation Docs}
     * 
     * Select which fields to find the minimum value
    **/
    _min?: SignersOnAssetPriceMinAggregateInputType
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/aggregations Aggregation Docs}
     * 
     * Select which fields to find the maximum value
    **/
    _max?: SignersOnAssetPriceMaxAggregateInputType
  }

  export type GetSignersOnAssetPriceAggregateType<T extends SignersOnAssetPriceAggregateArgs> = {
        [P in keyof T & keyof AggregateSignersOnAssetPrice]: P extends '_count' | 'count'
      ? T[P] extends true
        ? number
        : GetScalarType<T[P], AggregateSignersOnAssetPrice[P]>
      : GetScalarType<T[P], AggregateSignersOnAssetPrice[P]>
  }




  export type SignersOnAssetPriceGroupByArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    where?: SignersOnAssetPriceWhereInput
    orderBy?: SignersOnAssetPriceOrderByWithAggregationInput | SignersOnAssetPriceOrderByWithAggregationInput[]
    by: SignersOnAssetPriceScalarFieldEnum[] | SignersOnAssetPriceScalarFieldEnum
    having?: SignersOnAssetPriceScalarWhereWithAggregatesInput
    take?: number
    skip?: number
    _count?: SignersOnAssetPriceCountAggregateInputType | true
    _avg?: SignersOnAssetPriceAvgAggregateInputType
    _sum?: SignersOnAssetPriceSumAggregateInputType
    _min?: SignersOnAssetPriceMinAggregateInputType
    _max?: SignersOnAssetPriceMaxAggregateInputType
  }

  export type SignersOnAssetPriceGroupByOutputType = {
    signerId: number
    assetPriceId: number
    _count: SignersOnAssetPriceCountAggregateOutputType | null
    _avg: SignersOnAssetPriceAvgAggregateOutputType | null
    _sum: SignersOnAssetPriceSumAggregateOutputType | null
    _min: SignersOnAssetPriceMinAggregateOutputType | null
    _max: SignersOnAssetPriceMaxAggregateOutputType | null
  }

  type GetSignersOnAssetPriceGroupByPayload<T extends SignersOnAssetPriceGroupByArgs> = Prisma.PrismaPromise<
    Array<
      PickEnumerable<SignersOnAssetPriceGroupByOutputType, T['by']> &
        {
          [P in ((keyof T) & (keyof SignersOnAssetPriceGroupByOutputType))]: P extends '_count'
            ? T[P] extends boolean
              ? number
              : GetScalarType<T[P], SignersOnAssetPriceGroupByOutputType[P]>
            : GetScalarType<T[P], SignersOnAssetPriceGroupByOutputType[P]>
        }
      >
    >


  export type SignersOnAssetPriceSelect<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = $Extensions.GetSelect<{
    signerId?: boolean
    assetPriceId?: boolean
    assetPrice?: boolean | AssetPriceDefaultArgs<ExtArgs>
    signer?: boolean | SignerDefaultArgs<ExtArgs>
  }, ExtArgs["result"]["signersOnAssetPrice"]>

  export type SignersOnAssetPriceSelectScalar = {
    signerId?: boolean
    assetPriceId?: boolean
  }

  export type SignersOnAssetPriceInclude<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    assetPrice?: boolean | AssetPriceDefaultArgs<ExtArgs>
    signer?: boolean | SignerDefaultArgs<ExtArgs>
  }


  export type $SignersOnAssetPricePayload<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    name: "SignersOnAssetPrice"
    objects: {
      assetPrice: Prisma.$AssetPricePayload<ExtArgs>
      signer: Prisma.$SignerPayload<ExtArgs>
    }
    scalars: $Extensions.GetPayloadResult<{
      signerId: number
      assetPriceId: number
    }, ExtArgs["result"]["signersOnAssetPrice"]>
    composites: {}
  }


  type SignersOnAssetPriceGetPayload<S extends boolean | null | undefined | SignersOnAssetPriceDefaultArgs> = $Result.GetResult<Prisma.$SignersOnAssetPricePayload, S>

  type SignersOnAssetPriceCountArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = 
    Omit<SignersOnAssetPriceFindManyArgs, 'select' | 'include' | 'distinct' > & {
      select?: SignersOnAssetPriceCountAggregateInputType | true
    }

  export interface SignersOnAssetPriceDelegate<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> {
    [K: symbol]: { types: Prisma.TypeMap<ExtArgs>['model']['SignersOnAssetPrice'], meta: { name: 'SignersOnAssetPrice' } }
    /**
     * Find zero or one SignersOnAssetPrice that matches the filter.
     * @param {SignersOnAssetPriceFindUniqueArgs} args - Arguments to find a SignersOnAssetPrice
     * @example
     * // Get one SignersOnAssetPrice
     * const signersOnAssetPrice = await prisma.signersOnAssetPrice.findUnique({
     *   where: {
     *     // ... provide filter here
     *   }
     * })
    **/
    findUnique<T extends SignersOnAssetPriceFindUniqueArgs<ExtArgs>>(
      args: SelectSubset<T, SignersOnAssetPriceFindUniqueArgs<ExtArgs>>
    ): Prisma__SignersOnAssetPriceClient<$Result.GetResult<Prisma.$SignersOnAssetPricePayload<ExtArgs>, T, 'findUnique'> | null, null, ExtArgs>

    /**
     * Find one SignersOnAssetPrice that matches the filter or throw an error  with `error.code='P2025'` 
     *     if no matches were found.
     * @param {SignersOnAssetPriceFindUniqueOrThrowArgs} args - Arguments to find a SignersOnAssetPrice
     * @example
     * // Get one SignersOnAssetPrice
     * const signersOnAssetPrice = await prisma.signersOnAssetPrice.findUniqueOrThrow({
     *   where: {
     *     // ... provide filter here
     *   }
     * })
    **/
    findUniqueOrThrow<T extends SignersOnAssetPriceFindUniqueOrThrowArgs<ExtArgs>>(
      args?: SelectSubset<T, SignersOnAssetPriceFindUniqueOrThrowArgs<ExtArgs>>
    ): Prisma__SignersOnAssetPriceClient<$Result.GetResult<Prisma.$SignersOnAssetPricePayload<ExtArgs>, T, 'findUniqueOrThrow'>, never, ExtArgs>

    /**
     * Find the first SignersOnAssetPrice that matches the filter.
     * Note, that providing `undefined` is treated as the value not being there.
     * Read more here: https://pris.ly/d/null-undefined
     * @param {SignersOnAssetPriceFindFirstArgs} args - Arguments to find a SignersOnAssetPrice
     * @example
     * // Get one SignersOnAssetPrice
     * const signersOnAssetPrice = await prisma.signersOnAssetPrice.findFirst({
     *   where: {
     *     // ... provide filter here
     *   }
     * })
    **/
    findFirst<T extends SignersOnAssetPriceFindFirstArgs<ExtArgs>>(
      args?: SelectSubset<T, SignersOnAssetPriceFindFirstArgs<ExtArgs>>
    ): Prisma__SignersOnAssetPriceClient<$Result.GetResult<Prisma.$SignersOnAssetPricePayload<ExtArgs>, T, 'findFirst'> | null, null, ExtArgs>

    /**
     * Find the first SignersOnAssetPrice that matches the filter or
     * throw `PrismaKnownClientError` with `P2025` code if no matches were found.
     * Note, that providing `undefined` is treated as the value not being there.
     * Read more here: https://pris.ly/d/null-undefined
     * @param {SignersOnAssetPriceFindFirstOrThrowArgs} args - Arguments to find a SignersOnAssetPrice
     * @example
     * // Get one SignersOnAssetPrice
     * const signersOnAssetPrice = await prisma.signersOnAssetPrice.findFirstOrThrow({
     *   where: {
     *     // ... provide filter here
     *   }
     * })
    **/
    findFirstOrThrow<T extends SignersOnAssetPriceFindFirstOrThrowArgs<ExtArgs>>(
      args?: SelectSubset<T, SignersOnAssetPriceFindFirstOrThrowArgs<ExtArgs>>
    ): Prisma__SignersOnAssetPriceClient<$Result.GetResult<Prisma.$SignersOnAssetPricePayload<ExtArgs>, T, 'findFirstOrThrow'>, never, ExtArgs>

    /**
     * Find zero or more SignersOnAssetPrices that matches the filter.
     * Note, that providing `undefined` is treated as the value not being there.
     * Read more here: https://pris.ly/d/null-undefined
     * @param {SignersOnAssetPriceFindManyArgs=} args - Arguments to filter and select certain fields only.
     * @example
     * // Get all SignersOnAssetPrices
     * const signersOnAssetPrices = await prisma.signersOnAssetPrice.findMany()
     * 
     * // Get first 10 SignersOnAssetPrices
     * const signersOnAssetPrices = await prisma.signersOnAssetPrice.findMany({ take: 10 })
     * 
     * // Only select the `signerId`
     * const signersOnAssetPriceWithSignerIdOnly = await prisma.signersOnAssetPrice.findMany({ select: { signerId: true } })
     * 
    **/
    findMany<T extends SignersOnAssetPriceFindManyArgs<ExtArgs>>(
      args?: SelectSubset<T, SignersOnAssetPriceFindManyArgs<ExtArgs>>
    ): Prisma.PrismaPromise<$Result.GetResult<Prisma.$SignersOnAssetPricePayload<ExtArgs>, T, 'findMany'>>

    /**
     * Create a SignersOnAssetPrice.
     * @param {SignersOnAssetPriceCreateArgs} args - Arguments to create a SignersOnAssetPrice.
     * @example
     * // Create one SignersOnAssetPrice
     * const SignersOnAssetPrice = await prisma.signersOnAssetPrice.create({
     *   data: {
     *     // ... data to create a SignersOnAssetPrice
     *   }
     * })
     * 
    **/
    create<T extends SignersOnAssetPriceCreateArgs<ExtArgs>>(
      args: SelectSubset<T, SignersOnAssetPriceCreateArgs<ExtArgs>>
    ): Prisma__SignersOnAssetPriceClient<$Result.GetResult<Prisma.$SignersOnAssetPricePayload<ExtArgs>, T, 'create'>, never, ExtArgs>

    /**
     * Create many SignersOnAssetPrices.
     *     @param {SignersOnAssetPriceCreateManyArgs} args - Arguments to create many SignersOnAssetPrices.
     *     @example
     *     // Create many SignersOnAssetPrices
     *     const signersOnAssetPrice = await prisma.signersOnAssetPrice.createMany({
     *       data: {
     *         // ... provide data here
     *       }
     *     })
     *     
    **/
    createMany<T extends SignersOnAssetPriceCreateManyArgs<ExtArgs>>(
      args?: SelectSubset<T, SignersOnAssetPriceCreateManyArgs<ExtArgs>>
    ): Prisma.PrismaPromise<BatchPayload>

    /**
     * Delete a SignersOnAssetPrice.
     * @param {SignersOnAssetPriceDeleteArgs} args - Arguments to delete one SignersOnAssetPrice.
     * @example
     * // Delete one SignersOnAssetPrice
     * const SignersOnAssetPrice = await prisma.signersOnAssetPrice.delete({
     *   where: {
     *     // ... filter to delete one SignersOnAssetPrice
     *   }
     * })
     * 
    **/
    delete<T extends SignersOnAssetPriceDeleteArgs<ExtArgs>>(
      args: SelectSubset<T, SignersOnAssetPriceDeleteArgs<ExtArgs>>
    ): Prisma__SignersOnAssetPriceClient<$Result.GetResult<Prisma.$SignersOnAssetPricePayload<ExtArgs>, T, 'delete'>, never, ExtArgs>

    /**
     * Update one SignersOnAssetPrice.
     * @param {SignersOnAssetPriceUpdateArgs} args - Arguments to update one SignersOnAssetPrice.
     * @example
     * // Update one SignersOnAssetPrice
     * const signersOnAssetPrice = await prisma.signersOnAssetPrice.update({
     *   where: {
     *     // ... provide filter here
     *   },
     *   data: {
     *     // ... provide data here
     *   }
     * })
     * 
    **/
    update<T extends SignersOnAssetPriceUpdateArgs<ExtArgs>>(
      args: SelectSubset<T, SignersOnAssetPriceUpdateArgs<ExtArgs>>
    ): Prisma__SignersOnAssetPriceClient<$Result.GetResult<Prisma.$SignersOnAssetPricePayload<ExtArgs>, T, 'update'>, never, ExtArgs>

    /**
     * Delete zero or more SignersOnAssetPrices.
     * @param {SignersOnAssetPriceDeleteManyArgs} args - Arguments to filter SignersOnAssetPrices to delete.
     * @example
     * // Delete a few SignersOnAssetPrices
     * const { count } = await prisma.signersOnAssetPrice.deleteMany({
     *   where: {
     *     // ... provide filter here
     *   }
     * })
     * 
    **/
    deleteMany<T extends SignersOnAssetPriceDeleteManyArgs<ExtArgs>>(
      args?: SelectSubset<T, SignersOnAssetPriceDeleteManyArgs<ExtArgs>>
    ): Prisma.PrismaPromise<BatchPayload>

    /**
     * Update zero or more SignersOnAssetPrices.
     * Note, that providing `undefined` is treated as the value not being there.
     * Read more here: https://pris.ly/d/null-undefined
     * @param {SignersOnAssetPriceUpdateManyArgs} args - Arguments to update one or more rows.
     * @example
     * // Update many SignersOnAssetPrices
     * const signersOnAssetPrice = await prisma.signersOnAssetPrice.updateMany({
     *   where: {
     *     // ... provide filter here
     *   },
     *   data: {
     *     // ... provide data here
     *   }
     * })
     * 
    **/
    updateMany<T extends SignersOnAssetPriceUpdateManyArgs<ExtArgs>>(
      args: SelectSubset<T, SignersOnAssetPriceUpdateManyArgs<ExtArgs>>
    ): Prisma.PrismaPromise<BatchPayload>

    /**
     * Create or update one SignersOnAssetPrice.
     * @param {SignersOnAssetPriceUpsertArgs} args - Arguments to update or create a SignersOnAssetPrice.
     * @example
     * // Update or create a SignersOnAssetPrice
     * const signersOnAssetPrice = await prisma.signersOnAssetPrice.upsert({
     *   create: {
     *     // ... data to create a SignersOnAssetPrice
     *   },
     *   update: {
     *     // ... in case it already exists, update
     *   },
     *   where: {
     *     // ... the filter for the SignersOnAssetPrice we want to update
     *   }
     * })
    **/
    upsert<T extends SignersOnAssetPriceUpsertArgs<ExtArgs>>(
      args: SelectSubset<T, SignersOnAssetPriceUpsertArgs<ExtArgs>>
    ): Prisma__SignersOnAssetPriceClient<$Result.GetResult<Prisma.$SignersOnAssetPricePayload<ExtArgs>, T, 'upsert'>, never, ExtArgs>

    /**
     * Count the number of SignersOnAssetPrices.
     * Note, that providing `undefined` is treated as the value not being there.
     * Read more here: https://pris.ly/d/null-undefined
     * @param {SignersOnAssetPriceCountArgs} args - Arguments to filter SignersOnAssetPrices to count.
     * @example
     * // Count the number of SignersOnAssetPrices
     * const count = await prisma.signersOnAssetPrice.count({
     *   where: {
     *     // ... the filter for the SignersOnAssetPrices we want to count
     *   }
     * })
    **/
    count<T extends SignersOnAssetPriceCountArgs>(
      args?: Subset<T, SignersOnAssetPriceCountArgs>,
    ): Prisma.PrismaPromise<
      T extends $Utils.Record<'select', any>
        ? T['select'] extends true
          ? number
          : GetScalarType<T['select'], SignersOnAssetPriceCountAggregateOutputType>
        : number
    >

    /**
     * Allows you to perform aggregations operations on a SignersOnAssetPrice.
     * Note, that providing `undefined` is treated as the value not being there.
     * Read more here: https://pris.ly/d/null-undefined
     * @param {SignersOnAssetPriceAggregateArgs} args - Select which aggregations you would like to apply and on what fields.
     * @example
     * // Ordered by age ascending
     * // Where email contains prisma.io
     * // Limited to the 10 users
     * const aggregations = await prisma.user.aggregate({
     *   _avg: {
     *     age: true,
     *   },
     *   where: {
     *     email: {
     *       contains: "prisma.io",
     *     },
     *   },
     *   orderBy: {
     *     age: "asc",
     *   },
     *   take: 10,
     * })
    **/
    aggregate<T extends SignersOnAssetPriceAggregateArgs>(args: Subset<T, SignersOnAssetPriceAggregateArgs>): Prisma.PrismaPromise<GetSignersOnAssetPriceAggregateType<T>>

    /**
     * Group by SignersOnAssetPrice.
     * Note, that providing `undefined` is treated as the value not being there.
     * Read more here: https://pris.ly/d/null-undefined
     * @param {SignersOnAssetPriceGroupByArgs} args - Group by arguments.
     * @example
     * // Group by city, order by createdAt, get count
     * const result = await prisma.user.groupBy({
     *   by: ['city', 'createdAt'],
     *   orderBy: {
     *     createdAt: true
     *   },
     *   _count: {
     *     _all: true
     *   },
     * })
     * 
    **/
    groupBy<
      T extends SignersOnAssetPriceGroupByArgs,
      HasSelectOrTake extends Or<
        Extends<'skip', Keys<T>>,
        Extends<'take', Keys<T>>
      >,
      OrderByArg extends True extends HasSelectOrTake
        ? { orderBy: SignersOnAssetPriceGroupByArgs['orderBy'] }
        : { orderBy?: SignersOnAssetPriceGroupByArgs['orderBy'] },
      OrderFields extends ExcludeUnderscoreKeys<Keys<MaybeTupleToUnion<T['orderBy']>>>,
      ByFields extends MaybeTupleToUnion<T['by']>,
      ByValid extends Has<ByFields, OrderFields>,
      HavingFields extends GetHavingFields<T['having']>,
      HavingValid extends Has<ByFields, HavingFields>,
      ByEmpty extends T['by'] extends never[] ? True : False,
      InputErrors extends ByEmpty extends True
      ? `Error: "by" must not be empty.`
      : HavingValid extends False
      ? {
          [P in HavingFields]: P extends ByFields
            ? never
            : P extends string
            ? `Error: Field "${P}" used in "having" needs to be provided in "by".`
            : [
                Error,
                'Field ',
                P,
                ` in "having" needs to be provided in "by"`,
              ]
        }[HavingFields]
      : 'take' extends Keys<T>
      ? 'orderBy' extends Keys<T>
        ? ByValid extends True
          ? {}
          : {
              [P in OrderFields]: P extends ByFields
                ? never
                : `Error: Field "${P}" in "orderBy" needs to be provided in "by"`
            }[OrderFields]
        : 'Error: If you provide "take", you also need to provide "orderBy"'
      : 'skip' extends Keys<T>
      ? 'orderBy' extends Keys<T>
        ? ByValid extends True
          ? {}
          : {
              [P in OrderFields]: P extends ByFields
                ? never
                : `Error: Field "${P}" in "orderBy" needs to be provided in "by"`
            }[OrderFields]
        : 'Error: If you provide "skip", you also need to provide "orderBy"'
      : ByValid extends True
      ? {}
      : {
          [P in OrderFields]: P extends ByFields
            ? never
            : `Error: Field "${P}" in "orderBy" needs to be provided in "by"`
        }[OrderFields]
    >(args: SubsetIntersection<T, SignersOnAssetPriceGroupByArgs, OrderByArg> & InputErrors): {} extends InputErrors ? GetSignersOnAssetPriceGroupByPayload<T> : Prisma.PrismaPromise<InputErrors>
  /**
   * Fields of the SignersOnAssetPrice model
   */
  readonly fields: SignersOnAssetPriceFieldRefs;
  }

  /**
   * The delegate class that acts as a "Promise-like" for SignersOnAssetPrice.
   * Why is this prefixed with `Prisma__`?
   * Because we want to prevent naming conflicts as mentioned in
   * https://github.com/prisma/prisma-client-js/issues/707
   */
  export interface Prisma__SignersOnAssetPriceClient<T, Null = never, ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> extends Prisma.PrismaPromise<T> {
    readonly [Symbol.toStringTag]: 'PrismaPromise';

    assetPrice<T extends AssetPriceDefaultArgs<ExtArgs> = {}>(args?: Subset<T, AssetPriceDefaultArgs<ExtArgs>>): Prisma__AssetPriceClient<$Result.GetResult<Prisma.$AssetPricePayload<ExtArgs>, T, 'findUniqueOrThrow'> | Null, Null, ExtArgs>;

    signer<T extends SignerDefaultArgs<ExtArgs> = {}>(args?: Subset<T, SignerDefaultArgs<ExtArgs>>): Prisma__SignerClient<$Result.GetResult<Prisma.$SignerPayload<ExtArgs>, T, 'findUniqueOrThrow'> | Null, Null, ExtArgs>;

    /**
     * Attaches callbacks for the resolution and/or rejection of the Promise.
     * @param onfulfilled The callback to execute when the Promise is resolved.
     * @param onrejected The callback to execute when the Promise is rejected.
     * @returns A Promise for the completion of which ever callback is executed.
     */
    then<TResult1 = T, TResult2 = never>(onfulfilled?: ((value: T) => TResult1 | PromiseLike<TResult1>) | undefined | null, onrejected?: ((reason: any) => TResult2 | PromiseLike<TResult2>) | undefined | null): $Utils.JsPromise<TResult1 | TResult2>;
    /**
     * Attaches a callback for only the rejection of the Promise.
     * @param onrejected The callback to execute when the Promise is rejected.
     * @returns A Promise for the completion of the callback.
     */
    catch<TResult = never>(onrejected?: ((reason: any) => TResult | PromiseLike<TResult>) | undefined | null): $Utils.JsPromise<T | TResult>;
    /**
     * Attaches a callback that is invoked when the Promise is settled (fulfilled or rejected). The
     * resolved value cannot be modified from the callback.
     * @param onfinally The callback to execute when the Promise is settled (fulfilled or rejected).
     * @returns A Promise for the completion of the callback.
     */
    finally(onfinally?: (() => void) | undefined | null): $Utils.JsPromise<T>;
  }



  /**
   * Fields of the SignersOnAssetPrice model
   */ 
  interface SignersOnAssetPriceFieldRefs {
    readonly signerId: FieldRef<"SignersOnAssetPrice", 'Int'>
    readonly assetPriceId: FieldRef<"SignersOnAssetPrice", 'Int'>
  }
    

  // Custom InputTypes

  /**
   * SignersOnAssetPrice findUnique
   */
  export type SignersOnAssetPriceFindUniqueArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    /**
     * Select specific fields to fetch from the SignersOnAssetPrice
     */
    select?: SignersOnAssetPriceSelect<ExtArgs> | null
    /**
     * Choose, which related nodes to fetch as well.
     */
    include?: SignersOnAssetPriceInclude<ExtArgs> | null
    /**
     * Filter, which SignersOnAssetPrice to fetch.
     */
    where: SignersOnAssetPriceWhereUniqueInput
  }


  /**
   * SignersOnAssetPrice findUniqueOrThrow
   */
  export type SignersOnAssetPriceFindUniqueOrThrowArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    /**
     * Select specific fields to fetch from the SignersOnAssetPrice
     */
    select?: SignersOnAssetPriceSelect<ExtArgs> | null
    /**
     * Choose, which related nodes to fetch as well.
     */
    include?: SignersOnAssetPriceInclude<ExtArgs> | null
    /**
     * Filter, which SignersOnAssetPrice to fetch.
     */
    where: SignersOnAssetPriceWhereUniqueInput
  }


  /**
   * SignersOnAssetPrice findFirst
   */
  export type SignersOnAssetPriceFindFirstArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    /**
     * Select specific fields to fetch from the SignersOnAssetPrice
     */
    select?: SignersOnAssetPriceSelect<ExtArgs> | null
    /**
     * Choose, which related nodes to fetch as well.
     */
    include?: SignersOnAssetPriceInclude<ExtArgs> | null
    /**
     * Filter, which SignersOnAssetPrice to fetch.
     */
    where?: SignersOnAssetPriceWhereInput
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/sorting Sorting Docs}
     * 
     * Determine the order of SignersOnAssetPrices to fetch.
     */
    orderBy?: SignersOnAssetPriceOrderByWithRelationInput | SignersOnAssetPriceOrderByWithRelationInput[]
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/pagination#cursor-based-pagination Cursor Docs}
     * 
     * Sets the position for searching for SignersOnAssetPrices.
     */
    cursor?: SignersOnAssetPriceWhereUniqueInput
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/pagination Pagination Docs}
     * 
     * Take `±n` SignersOnAssetPrices from the position of the cursor.
     */
    take?: number
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/pagination Pagination Docs}
     * 
     * Skip the first `n` SignersOnAssetPrices.
     */
    skip?: number
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/distinct Distinct Docs}
     * 
     * Filter by unique combinations of SignersOnAssetPrices.
     */
    distinct?: SignersOnAssetPriceScalarFieldEnum | SignersOnAssetPriceScalarFieldEnum[]
  }


  /**
   * SignersOnAssetPrice findFirstOrThrow
   */
  export type SignersOnAssetPriceFindFirstOrThrowArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    /**
     * Select specific fields to fetch from the SignersOnAssetPrice
     */
    select?: SignersOnAssetPriceSelect<ExtArgs> | null
    /**
     * Choose, which related nodes to fetch as well.
     */
    include?: SignersOnAssetPriceInclude<ExtArgs> | null
    /**
     * Filter, which SignersOnAssetPrice to fetch.
     */
    where?: SignersOnAssetPriceWhereInput
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/sorting Sorting Docs}
     * 
     * Determine the order of SignersOnAssetPrices to fetch.
     */
    orderBy?: SignersOnAssetPriceOrderByWithRelationInput | SignersOnAssetPriceOrderByWithRelationInput[]
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/pagination#cursor-based-pagination Cursor Docs}
     * 
     * Sets the position for searching for SignersOnAssetPrices.
     */
    cursor?: SignersOnAssetPriceWhereUniqueInput
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/pagination Pagination Docs}
     * 
     * Take `±n` SignersOnAssetPrices from the position of the cursor.
     */
    take?: number
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/pagination Pagination Docs}
     * 
     * Skip the first `n` SignersOnAssetPrices.
     */
    skip?: number
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/distinct Distinct Docs}
     * 
     * Filter by unique combinations of SignersOnAssetPrices.
     */
    distinct?: SignersOnAssetPriceScalarFieldEnum | SignersOnAssetPriceScalarFieldEnum[]
  }


  /**
   * SignersOnAssetPrice findMany
   */
  export type SignersOnAssetPriceFindManyArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    /**
     * Select specific fields to fetch from the SignersOnAssetPrice
     */
    select?: SignersOnAssetPriceSelect<ExtArgs> | null
    /**
     * Choose, which related nodes to fetch as well.
     */
    include?: SignersOnAssetPriceInclude<ExtArgs> | null
    /**
     * Filter, which SignersOnAssetPrices to fetch.
     */
    where?: SignersOnAssetPriceWhereInput
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/sorting Sorting Docs}
     * 
     * Determine the order of SignersOnAssetPrices to fetch.
     */
    orderBy?: SignersOnAssetPriceOrderByWithRelationInput | SignersOnAssetPriceOrderByWithRelationInput[]
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/pagination#cursor-based-pagination Cursor Docs}
     * 
     * Sets the position for listing SignersOnAssetPrices.
     */
    cursor?: SignersOnAssetPriceWhereUniqueInput
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/pagination Pagination Docs}
     * 
     * Take `±n` SignersOnAssetPrices from the position of the cursor.
     */
    take?: number
    /**
     * {@link https://www.prisma.io/docs/concepts/components/prisma-client/pagination Pagination Docs}
     * 
     * Skip the first `n` SignersOnAssetPrices.
     */
    skip?: number
    distinct?: SignersOnAssetPriceScalarFieldEnum | SignersOnAssetPriceScalarFieldEnum[]
  }


  /**
   * SignersOnAssetPrice create
   */
  export type SignersOnAssetPriceCreateArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    /**
     * Select specific fields to fetch from the SignersOnAssetPrice
     */
    select?: SignersOnAssetPriceSelect<ExtArgs> | null
    /**
     * Choose, which related nodes to fetch as well.
     */
    include?: SignersOnAssetPriceInclude<ExtArgs> | null
    /**
     * The data needed to create a SignersOnAssetPrice.
     */
    data: XOR<SignersOnAssetPriceCreateInput, SignersOnAssetPriceUncheckedCreateInput>
  }


  /**
   * SignersOnAssetPrice createMany
   */
  export type SignersOnAssetPriceCreateManyArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    /**
     * The data used to create many SignersOnAssetPrices.
     */
    data: SignersOnAssetPriceCreateManyInput | SignersOnAssetPriceCreateManyInput[]
    skipDuplicates?: boolean
  }


  /**
   * SignersOnAssetPrice update
   */
  export type SignersOnAssetPriceUpdateArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    /**
     * Select specific fields to fetch from the SignersOnAssetPrice
     */
    select?: SignersOnAssetPriceSelect<ExtArgs> | null
    /**
     * Choose, which related nodes to fetch as well.
     */
    include?: SignersOnAssetPriceInclude<ExtArgs> | null
    /**
     * The data needed to update a SignersOnAssetPrice.
     */
    data: XOR<SignersOnAssetPriceUpdateInput, SignersOnAssetPriceUncheckedUpdateInput>
    /**
     * Choose, which SignersOnAssetPrice to update.
     */
    where: SignersOnAssetPriceWhereUniqueInput
  }


  /**
   * SignersOnAssetPrice updateMany
   */
  export type SignersOnAssetPriceUpdateManyArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    /**
     * The data used to update SignersOnAssetPrices.
     */
    data: XOR<SignersOnAssetPriceUpdateManyMutationInput, SignersOnAssetPriceUncheckedUpdateManyInput>
    /**
     * Filter which SignersOnAssetPrices to update
     */
    where?: SignersOnAssetPriceWhereInput
  }


  /**
   * SignersOnAssetPrice upsert
   */
  export type SignersOnAssetPriceUpsertArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    /**
     * Select specific fields to fetch from the SignersOnAssetPrice
     */
    select?: SignersOnAssetPriceSelect<ExtArgs> | null
    /**
     * Choose, which related nodes to fetch as well.
     */
    include?: SignersOnAssetPriceInclude<ExtArgs> | null
    /**
     * The filter to search for the SignersOnAssetPrice to update in case it exists.
     */
    where: SignersOnAssetPriceWhereUniqueInput
    /**
     * In case the SignersOnAssetPrice found by the `where` argument doesn't exist, create a new SignersOnAssetPrice with this data.
     */
    create: XOR<SignersOnAssetPriceCreateInput, SignersOnAssetPriceUncheckedCreateInput>
    /**
     * In case the SignersOnAssetPrice was found with the provided `where` argument, update it with this data.
     */
    update: XOR<SignersOnAssetPriceUpdateInput, SignersOnAssetPriceUncheckedUpdateInput>
  }


  /**
   * SignersOnAssetPrice delete
   */
  export type SignersOnAssetPriceDeleteArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    /**
     * Select specific fields to fetch from the SignersOnAssetPrice
     */
    select?: SignersOnAssetPriceSelect<ExtArgs> | null
    /**
     * Choose, which related nodes to fetch as well.
     */
    include?: SignersOnAssetPriceInclude<ExtArgs> | null
    /**
     * Filter which SignersOnAssetPrice to delete.
     */
    where: SignersOnAssetPriceWhereUniqueInput
  }


  /**
   * SignersOnAssetPrice deleteMany
   */
  export type SignersOnAssetPriceDeleteManyArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    /**
     * Filter which SignersOnAssetPrices to delete
     */
    where?: SignersOnAssetPriceWhereInput
  }


  /**
   * SignersOnAssetPrice without action
   */
  export type SignersOnAssetPriceDefaultArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = {
    /**
     * Select specific fields to fetch from the SignersOnAssetPrice
     */
    select?: SignersOnAssetPriceSelect<ExtArgs> | null
    /**
     * Choose, which related nodes to fetch as well.
     */
    include?: SignersOnAssetPriceInclude<ExtArgs> | null
  }



  /**
   * Enums
   */

  export const TransactionIsolationLevel: {
    ReadUncommitted: 'ReadUncommitted',
    ReadCommitted: 'ReadCommitted',
    RepeatableRead: 'RepeatableRead',
    Serializable: 'Serializable'
  };

  export type TransactionIsolationLevel = (typeof TransactionIsolationLevel)[keyof typeof TransactionIsolationLevel]


  export const AssetPriceScalarFieldEnum: {
    id: 'id',
    dataSetId: 'dataSetId',
    createdAt: 'createdAt',
    updatedAt: 'updatedAt',
    block: 'block',
    price: 'price',
    signature: 'signature'
  };

  export type AssetPriceScalarFieldEnum = (typeof AssetPriceScalarFieldEnum)[keyof typeof AssetPriceScalarFieldEnum]


  export const DataSetScalarFieldEnum: {
    id: 'id',
    name: 'name'
  };

  export type DataSetScalarFieldEnum = (typeof DataSetScalarFieldEnum)[keyof typeof DataSetScalarFieldEnum]


  export const SignerScalarFieldEnum: {
    id: 'id',
    key: 'key',
    name: 'name'
  };

  export type SignerScalarFieldEnum = (typeof SignerScalarFieldEnum)[keyof typeof SignerScalarFieldEnum]


  export const SignersOnAssetPriceScalarFieldEnum: {
    signerId: 'signerId',
    assetPriceId: 'assetPriceId'
  };

  export type SignersOnAssetPriceScalarFieldEnum = (typeof SignersOnAssetPriceScalarFieldEnum)[keyof typeof SignersOnAssetPriceScalarFieldEnum]


  export const SortOrder: {
    asc: 'asc',
    desc: 'desc'
  };

  export type SortOrder = (typeof SortOrder)[keyof typeof SortOrder]


  export const QueryMode: {
    default: 'default',
    insensitive: 'insensitive'
  };

  export type QueryMode = (typeof QueryMode)[keyof typeof QueryMode]


  export const NullsOrder: {
    first: 'first',
    last: 'last'
  };

  export type NullsOrder = (typeof NullsOrder)[keyof typeof NullsOrder]


  /**
   * Field references 
   */


  /**
   * Reference to a field of type 'Int'
   */
  export type IntFieldRefInput<$PrismaModel> = FieldRefInputType<$PrismaModel, 'Int'>
    


  /**
   * Reference to a field of type 'Int[]'
   */
  export type ListIntFieldRefInput<$PrismaModel> = FieldRefInputType<$PrismaModel, 'Int[]'>
    


  /**
   * Reference to a field of type 'DateTime'
   */
  export type DateTimeFieldRefInput<$PrismaModel> = FieldRefInputType<$PrismaModel, 'DateTime'>
    


  /**
   * Reference to a field of type 'DateTime[]'
   */
  export type ListDateTimeFieldRefInput<$PrismaModel> = FieldRefInputType<$PrismaModel, 'DateTime[]'>
    


  /**
   * Reference to a field of type 'Decimal'
   */
  export type DecimalFieldRefInput<$PrismaModel> = FieldRefInputType<$PrismaModel, 'Decimal'>
    


  /**
   * Reference to a field of type 'Decimal[]'
   */
  export type ListDecimalFieldRefInput<$PrismaModel> = FieldRefInputType<$PrismaModel, 'Decimal[]'>
    


  /**
   * Reference to a field of type 'String'
   */
  export type StringFieldRefInput<$PrismaModel> = FieldRefInputType<$PrismaModel, 'String'>
    


  /**
   * Reference to a field of type 'String[]'
   */
  export type ListStringFieldRefInput<$PrismaModel> = FieldRefInputType<$PrismaModel, 'String[]'>
    


  /**
   * Reference to a field of type 'Float'
   */
  export type FloatFieldRefInput<$PrismaModel> = FieldRefInputType<$PrismaModel, 'Float'>
    


  /**
   * Reference to a field of type 'Float[]'
   */
  export type ListFloatFieldRefInput<$PrismaModel> = FieldRefInputType<$PrismaModel, 'Float[]'>
    
  /**
   * Deep Input Types
   */


  export type AssetPriceWhereInput = {
    AND?: AssetPriceWhereInput | AssetPriceWhereInput[]
    OR?: AssetPriceWhereInput[]
    NOT?: AssetPriceWhereInput | AssetPriceWhereInput[]
    id?: IntFilter<"AssetPrice"> | number
    dataSetId?: IntFilter<"AssetPrice"> | number
    createdAt?: DateTimeFilter<"AssetPrice"> | Date | string
    updatedAt?: DateTimeFilter<"AssetPrice"> | Date | string
    block?: IntNullableFilter<"AssetPrice"> | number | null
    price?: DecimalFilter<"AssetPrice"> | Decimal | DecimalJsLike | number | string
    signature?: StringFilter<"AssetPrice"> | string
    dataset?: XOR<DataSetRelationFilter, DataSetWhereInput>
    signersOnAssetPrice?: SignersOnAssetPriceListRelationFilter
  }

  export type AssetPriceOrderByWithRelationInput = {
    id?: SortOrder
    dataSetId?: SortOrder
    createdAt?: SortOrder
    updatedAt?: SortOrder
    block?: SortOrderInput | SortOrder
    price?: SortOrder
    signature?: SortOrder
    dataset?: DataSetOrderByWithRelationInput
    signersOnAssetPrice?: SignersOnAssetPriceOrderByRelationAggregateInput
  }

  export type AssetPriceWhereUniqueInput = Prisma.AtLeast<{
    id?: number
    dataSetId_block?: AssetPriceDataSetIdBlockCompoundUniqueInput
    AND?: AssetPriceWhereInput | AssetPriceWhereInput[]
    OR?: AssetPriceWhereInput[]
    NOT?: AssetPriceWhereInput | AssetPriceWhereInput[]
    dataSetId?: IntFilter<"AssetPrice"> | number
    createdAt?: DateTimeFilter<"AssetPrice"> | Date | string
    updatedAt?: DateTimeFilter<"AssetPrice"> | Date | string
    block?: IntNullableFilter<"AssetPrice"> | number | null
    price?: DecimalFilter<"AssetPrice"> | Decimal | DecimalJsLike | number | string
    signature?: StringFilter<"AssetPrice"> | string
    dataset?: XOR<DataSetRelationFilter, DataSetWhereInput>
    signersOnAssetPrice?: SignersOnAssetPriceListRelationFilter
  }, "id" | "dataSetId_block">

  export type AssetPriceOrderByWithAggregationInput = {
    id?: SortOrder
    dataSetId?: SortOrder
    createdAt?: SortOrder
    updatedAt?: SortOrder
    block?: SortOrderInput | SortOrder
    price?: SortOrder
    signature?: SortOrder
    _count?: AssetPriceCountOrderByAggregateInput
    _avg?: AssetPriceAvgOrderByAggregateInput
    _max?: AssetPriceMaxOrderByAggregateInput
    _min?: AssetPriceMinOrderByAggregateInput
    _sum?: AssetPriceSumOrderByAggregateInput
  }

  export type AssetPriceScalarWhereWithAggregatesInput = {
    AND?: AssetPriceScalarWhereWithAggregatesInput | AssetPriceScalarWhereWithAggregatesInput[]
    OR?: AssetPriceScalarWhereWithAggregatesInput[]
    NOT?: AssetPriceScalarWhereWithAggregatesInput | AssetPriceScalarWhereWithAggregatesInput[]
    id?: IntWithAggregatesFilter<"AssetPrice"> | number
    dataSetId?: IntWithAggregatesFilter<"AssetPrice"> | number
    createdAt?: DateTimeWithAggregatesFilter<"AssetPrice"> | Date | string
    updatedAt?: DateTimeWithAggregatesFilter<"AssetPrice"> | Date | string
    block?: IntNullableWithAggregatesFilter<"AssetPrice"> | number | null
    price?: DecimalWithAggregatesFilter<"AssetPrice"> | Decimal | DecimalJsLike | number | string
    signature?: StringWithAggregatesFilter<"AssetPrice"> | string
  }

  export type DataSetWhereInput = {
    AND?: DataSetWhereInput | DataSetWhereInput[]
    OR?: DataSetWhereInput[]
    NOT?: DataSetWhereInput | DataSetWhereInput[]
    id?: IntFilter<"DataSet"> | number
    name?: StringFilter<"DataSet"> | string
    AssetPrice?: AssetPriceListRelationFilter
  }

  export type DataSetOrderByWithRelationInput = {
    id?: SortOrder
    name?: SortOrder
    AssetPrice?: AssetPriceOrderByRelationAggregateInput
  }

  export type DataSetWhereUniqueInput = Prisma.AtLeast<{
    id?: number
    name?: string
    AND?: DataSetWhereInput | DataSetWhereInput[]
    OR?: DataSetWhereInput[]
    NOT?: DataSetWhereInput | DataSetWhereInput[]
    AssetPrice?: AssetPriceListRelationFilter
  }, "id" | "name">

  export type DataSetOrderByWithAggregationInput = {
    id?: SortOrder
    name?: SortOrder
    _count?: DataSetCountOrderByAggregateInput
    _avg?: DataSetAvgOrderByAggregateInput
    _max?: DataSetMaxOrderByAggregateInput
    _min?: DataSetMinOrderByAggregateInput
    _sum?: DataSetSumOrderByAggregateInput
  }

  export type DataSetScalarWhereWithAggregatesInput = {
    AND?: DataSetScalarWhereWithAggregatesInput | DataSetScalarWhereWithAggregatesInput[]
    OR?: DataSetScalarWhereWithAggregatesInput[]
    NOT?: DataSetScalarWhereWithAggregatesInput | DataSetScalarWhereWithAggregatesInput[]
    id?: IntWithAggregatesFilter<"DataSet"> | number
    name?: StringWithAggregatesFilter<"DataSet"> | string
  }

  export type SignerWhereInput = {
    AND?: SignerWhereInput | SignerWhereInput[]
    OR?: SignerWhereInput[]
    NOT?: SignerWhereInput | SignerWhereInput[]
    id?: IntFilter<"Signer"> | number
    key?: StringFilter<"Signer"> | string
    name?: StringNullableFilter<"Signer"> | string | null
    signersOnAssetPrice?: SignersOnAssetPriceListRelationFilter
  }

  export type SignerOrderByWithRelationInput = {
    id?: SortOrder
    key?: SortOrder
    name?: SortOrderInput | SortOrder
    signersOnAssetPrice?: SignersOnAssetPriceOrderByRelationAggregateInput
  }

  export type SignerWhereUniqueInput = Prisma.AtLeast<{
    id?: number
    key?: string
    AND?: SignerWhereInput | SignerWhereInput[]
    OR?: SignerWhereInput[]
    NOT?: SignerWhereInput | SignerWhereInput[]
    name?: StringNullableFilter<"Signer"> | string | null
    signersOnAssetPrice?: SignersOnAssetPriceListRelationFilter
  }, "id" | "key">

  export type SignerOrderByWithAggregationInput = {
    id?: SortOrder
    key?: SortOrder
    name?: SortOrderInput | SortOrder
    _count?: SignerCountOrderByAggregateInput
    _avg?: SignerAvgOrderByAggregateInput
    _max?: SignerMaxOrderByAggregateInput
    _min?: SignerMinOrderByAggregateInput
    _sum?: SignerSumOrderByAggregateInput
  }

  export type SignerScalarWhereWithAggregatesInput = {
    AND?: SignerScalarWhereWithAggregatesInput | SignerScalarWhereWithAggregatesInput[]
    OR?: SignerScalarWhereWithAggregatesInput[]
    NOT?: SignerScalarWhereWithAggregatesInput | SignerScalarWhereWithAggregatesInput[]
    id?: IntWithAggregatesFilter<"Signer"> | number
    key?: StringWithAggregatesFilter<"Signer"> | string
    name?: StringNullableWithAggregatesFilter<"Signer"> | string | null
  }

  export type SignersOnAssetPriceWhereInput = {
    AND?: SignersOnAssetPriceWhereInput | SignersOnAssetPriceWhereInput[]
    OR?: SignersOnAssetPriceWhereInput[]
    NOT?: SignersOnAssetPriceWhereInput | SignersOnAssetPriceWhereInput[]
    signerId?: IntFilter<"SignersOnAssetPrice"> | number
    assetPriceId?: IntFilter<"SignersOnAssetPrice"> | number
    assetPrice?: XOR<AssetPriceRelationFilter, AssetPriceWhereInput>
    signer?: XOR<SignerRelationFilter, SignerWhereInput>
  }

  export type SignersOnAssetPriceOrderByWithRelationInput = {
    signerId?: SortOrder
    assetPriceId?: SortOrder
    assetPrice?: AssetPriceOrderByWithRelationInput
    signer?: SignerOrderByWithRelationInput
  }

  export type SignersOnAssetPriceWhereUniqueInput = Prisma.AtLeast<{
    signerId_assetPriceId?: SignersOnAssetPriceSignerIdAssetPriceIdCompoundUniqueInput
    AND?: SignersOnAssetPriceWhereInput | SignersOnAssetPriceWhereInput[]
    OR?: SignersOnAssetPriceWhereInput[]
    NOT?: SignersOnAssetPriceWhereInput | SignersOnAssetPriceWhereInput[]
    signerId?: IntFilter<"SignersOnAssetPrice"> | number
    assetPriceId?: IntFilter<"SignersOnAssetPrice"> | number
    assetPrice?: XOR<AssetPriceRelationFilter, AssetPriceWhereInput>
    signer?: XOR<SignerRelationFilter, SignerWhereInput>
  }, "signerId_assetPriceId">

  export type SignersOnAssetPriceOrderByWithAggregationInput = {
    signerId?: SortOrder
    assetPriceId?: SortOrder
    _count?: SignersOnAssetPriceCountOrderByAggregateInput
    _avg?: SignersOnAssetPriceAvgOrderByAggregateInput
    _max?: SignersOnAssetPriceMaxOrderByAggregateInput
    _min?: SignersOnAssetPriceMinOrderByAggregateInput
    _sum?: SignersOnAssetPriceSumOrderByAggregateInput
  }

  export type SignersOnAssetPriceScalarWhereWithAggregatesInput = {
    AND?: SignersOnAssetPriceScalarWhereWithAggregatesInput | SignersOnAssetPriceScalarWhereWithAggregatesInput[]
    OR?: SignersOnAssetPriceScalarWhereWithAggregatesInput[]
    NOT?: SignersOnAssetPriceScalarWhereWithAggregatesInput | SignersOnAssetPriceScalarWhereWithAggregatesInput[]
    signerId?: IntWithAggregatesFilter<"SignersOnAssetPrice"> | number
    assetPriceId?: IntWithAggregatesFilter<"SignersOnAssetPrice"> | number
  }

  export type AssetPriceCreateInput = {
    createdAt?: Date | string
    updatedAt?: Date | string
    block?: number | null
    price: Decimal | DecimalJsLike | number | string
    signature: string
    dataset: DataSetCreateNestedOneWithoutAssetPriceInput
    signersOnAssetPrice?: SignersOnAssetPriceCreateNestedManyWithoutAssetPriceInput
  }

  export type AssetPriceUncheckedCreateInput = {
    id?: number
    dataSetId: number
    createdAt?: Date | string
    updatedAt?: Date | string
    block?: number | null
    price: Decimal | DecimalJsLike | number | string
    signature: string
    signersOnAssetPrice?: SignersOnAssetPriceUncheckedCreateNestedManyWithoutAssetPriceInput
  }

  export type AssetPriceUpdateInput = {
    createdAt?: DateTimeFieldUpdateOperationsInput | Date | string
    updatedAt?: DateTimeFieldUpdateOperationsInput | Date | string
    block?: NullableIntFieldUpdateOperationsInput | number | null
    price?: DecimalFieldUpdateOperationsInput | Decimal | DecimalJsLike | number | string
    signature?: StringFieldUpdateOperationsInput | string
    dataset?: DataSetUpdateOneRequiredWithoutAssetPriceNestedInput
    signersOnAssetPrice?: SignersOnAssetPriceUpdateManyWithoutAssetPriceNestedInput
  }

  export type AssetPriceUncheckedUpdateInput = {
    id?: IntFieldUpdateOperationsInput | number
    dataSetId?: IntFieldUpdateOperationsInput | number
    createdAt?: DateTimeFieldUpdateOperationsInput | Date | string
    updatedAt?: DateTimeFieldUpdateOperationsInput | Date | string
    block?: NullableIntFieldUpdateOperationsInput | number | null
    price?: DecimalFieldUpdateOperationsInput | Decimal | DecimalJsLike | number | string
    signature?: StringFieldUpdateOperationsInput | string
    signersOnAssetPrice?: SignersOnAssetPriceUncheckedUpdateManyWithoutAssetPriceNestedInput
  }

  export type AssetPriceCreateManyInput = {
    id?: number
    dataSetId: number
    createdAt?: Date | string
    updatedAt?: Date | string
    block?: number | null
    price: Decimal | DecimalJsLike | number | string
    signature: string
  }

  export type AssetPriceUpdateManyMutationInput = {
    createdAt?: DateTimeFieldUpdateOperationsInput | Date | string
    updatedAt?: DateTimeFieldUpdateOperationsInput | Date | string
    block?: NullableIntFieldUpdateOperationsInput | number | null
    price?: DecimalFieldUpdateOperationsInput | Decimal | DecimalJsLike | number | string
    signature?: StringFieldUpdateOperationsInput | string
  }

  export type AssetPriceUncheckedUpdateManyInput = {
    id?: IntFieldUpdateOperationsInput | number
    dataSetId?: IntFieldUpdateOperationsInput | number
    createdAt?: DateTimeFieldUpdateOperationsInput | Date | string
    updatedAt?: DateTimeFieldUpdateOperationsInput | Date | string
    block?: NullableIntFieldUpdateOperationsInput | number | null
    price?: DecimalFieldUpdateOperationsInput | Decimal | DecimalJsLike | number | string
    signature?: StringFieldUpdateOperationsInput | string
  }

  export type DataSetCreateInput = {
    name: string
    AssetPrice?: AssetPriceCreateNestedManyWithoutDatasetInput
  }

  export type DataSetUncheckedCreateInput = {
    id?: number
    name: string
    AssetPrice?: AssetPriceUncheckedCreateNestedManyWithoutDatasetInput
  }

  export type DataSetUpdateInput = {
    name?: StringFieldUpdateOperationsInput | string
    AssetPrice?: AssetPriceUpdateManyWithoutDatasetNestedInput
  }

  export type DataSetUncheckedUpdateInput = {
    id?: IntFieldUpdateOperationsInput | number
    name?: StringFieldUpdateOperationsInput | string
    AssetPrice?: AssetPriceUncheckedUpdateManyWithoutDatasetNestedInput
  }

  export type DataSetCreateManyInput = {
    id?: number
    name: string
  }

  export type DataSetUpdateManyMutationInput = {
    name?: StringFieldUpdateOperationsInput | string
  }

  export type DataSetUncheckedUpdateManyInput = {
    id?: IntFieldUpdateOperationsInput | number
    name?: StringFieldUpdateOperationsInput | string
  }

  export type SignerCreateInput = {
    key: string
    name?: string | null
    signersOnAssetPrice?: SignersOnAssetPriceCreateNestedManyWithoutSignerInput
  }

  export type SignerUncheckedCreateInput = {
    id?: number
    key: string
    name?: string | null
    signersOnAssetPrice?: SignersOnAssetPriceUncheckedCreateNestedManyWithoutSignerInput
  }

  export type SignerUpdateInput = {
    key?: StringFieldUpdateOperationsInput | string
    name?: NullableStringFieldUpdateOperationsInput | string | null
    signersOnAssetPrice?: SignersOnAssetPriceUpdateManyWithoutSignerNestedInput
  }

  export type SignerUncheckedUpdateInput = {
    id?: IntFieldUpdateOperationsInput | number
    key?: StringFieldUpdateOperationsInput | string
    name?: NullableStringFieldUpdateOperationsInput | string | null
    signersOnAssetPrice?: SignersOnAssetPriceUncheckedUpdateManyWithoutSignerNestedInput
  }

  export type SignerCreateManyInput = {
    id?: number
    key: string
    name?: string | null
  }

  export type SignerUpdateManyMutationInput = {
    key?: StringFieldUpdateOperationsInput | string
    name?: NullableStringFieldUpdateOperationsInput | string | null
  }

  export type SignerUncheckedUpdateManyInput = {
    id?: IntFieldUpdateOperationsInput | number
    key?: StringFieldUpdateOperationsInput | string
    name?: NullableStringFieldUpdateOperationsInput | string | null
  }

  export type SignersOnAssetPriceCreateInput = {
    assetPrice: AssetPriceCreateNestedOneWithoutSignersOnAssetPriceInput
    signer: SignerCreateNestedOneWithoutSignersOnAssetPriceInput
  }

  export type SignersOnAssetPriceUncheckedCreateInput = {
    signerId: number
    assetPriceId: number
  }

  export type SignersOnAssetPriceUpdateInput = {
    assetPrice?: AssetPriceUpdateOneRequiredWithoutSignersOnAssetPriceNestedInput
    signer?: SignerUpdateOneRequiredWithoutSignersOnAssetPriceNestedInput
  }

  export type SignersOnAssetPriceUncheckedUpdateInput = {
    signerId?: IntFieldUpdateOperationsInput | number
    assetPriceId?: IntFieldUpdateOperationsInput | number
  }

  export type SignersOnAssetPriceCreateManyInput = {
    signerId: number
    assetPriceId: number
  }

  export type SignersOnAssetPriceUpdateManyMutationInput = {

  }

  export type SignersOnAssetPriceUncheckedUpdateManyInput = {
    signerId?: IntFieldUpdateOperationsInput | number
    assetPriceId?: IntFieldUpdateOperationsInput | number
  }

  export type IntFilter<$PrismaModel = never> = {
    equals?: number | IntFieldRefInput<$PrismaModel>
    in?: number[] | ListIntFieldRefInput<$PrismaModel>
    notIn?: number[] | ListIntFieldRefInput<$PrismaModel>
    lt?: number | IntFieldRefInput<$PrismaModel>
    lte?: number | IntFieldRefInput<$PrismaModel>
    gt?: number | IntFieldRefInput<$PrismaModel>
    gte?: number | IntFieldRefInput<$PrismaModel>
    not?: NestedIntFilter<$PrismaModel> | number
  }

  export type DateTimeFilter<$PrismaModel = never> = {
    equals?: Date | string | DateTimeFieldRefInput<$PrismaModel>
    in?: Date[] | string[] | ListDateTimeFieldRefInput<$PrismaModel>
    notIn?: Date[] | string[] | ListDateTimeFieldRefInput<$PrismaModel>
    lt?: Date | string | DateTimeFieldRefInput<$PrismaModel>
    lte?: Date | string | DateTimeFieldRefInput<$PrismaModel>
    gt?: Date | string | DateTimeFieldRefInput<$PrismaModel>
    gte?: Date | string | DateTimeFieldRefInput<$PrismaModel>
    not?: NestedDateTimeFilter<$PrismaModel> | Date | string
  }

  export type IntNullableFilter<$PrismaModel = never> = {
    equals?: number | IntFieldRefInput<$PrismaModel> | null
    in?: number[] | ListIntFieldRefInput<$PrismaModel> | null
    notIn?: number[] | ListIntFieldRefInput<$PrismaModel> | null
    lt?: number | IntFieldRefInput<$PrismaModel>
    lte?: number | IntFieldRefInput<$PrismaModel>
    gt?: number | IntFieldRefInput<$PrismaModel>
    gte?: number | IntFieldRefInput<$PrismaModel>
    not?: NestedIntNullableFilter<$PrismaModel> | number | null
  }

  export type DecimalFilter<$PrismaModel = never> = {
    equals?: Decimal | DecimalJsLike | number | string | DecimalFieldRefInput<$PrismaModel>
    in?: Decimal[] | DecimalJsLike[] | number[] | string[] | ListDecimalFieldRefInput<$PrismaModel>
    notIn?: Decimal[] | DecimalJsLike[] | number[] | string[] | ListDecimalFieldRefInput<$PrismaModel>
    lt?: Decimal | DecimalJsLike | number | string | DecimalFieldRefInput<$PrismaModel>
    lte?: Decimal | DecimalJsLike | number | string | DecimalFieldRefInput<$PrismaModel>
    gt?: Decimal | DecimalJsLike | number | string | DecimalFieldRefInput<$PrismaModel>
    gte?: Decimal | DecimalJsLike | number | string | DecimalFieldRefInput<$PrismaModel>
    not?: NestedDecimalFilter<$PrismaModel> | Decimal | DecimalJsLike | number | string
  }

  export type StringFilter<$PrismaModel = never> = {
    equals?: string | StringFieldRefInput<$PrismaModel>
    in?: string[] | ListStringFieldRefInput<$PrismaModel>
    notIn?: string[] | ListStringFieldRefInput<$PrismaModel>
    lt?: string | StringFieldRefInput<$PrismaModel>
    lte?: string | StringFieldRefInput<$PrismaModel>
    gt?: string | StringFieldRefInput<$PrismaModel>
    gte?: string | StringFieldRefInput<$PrismaModel>
    contains?: string | StringFieldRefInput<$PrismaModel>
    startsWith?: string | StringFieldRefInput<$PrismaModel>
    endsWith?: string | StringFieldRefInput<$PrismaModel>
    mode?: QueryMode
    not?: NestedStringFilter<$PrismaModel> | string
  }

  export type DataSetRelationFilter = {
    is?: DataSetWhereInput
    isNot?: DataSetWhereInput
  }

  export type SignersOnAssetPriceListRelationFilter = {
    every?: SignersOnAssetPriceWhereInput
    some?: SignersOnAssetPriceWhereInput
    none?: SignersOnAssetPriceWhereInput
  }

  export type SortOrderInput = {
    sort: SortOrder
    nulls?: NullsOrder
  }

  export type SignersOnAssetPriceOrderByRelationAggregateInput = {
    _count?: SortOrder
  }

  export type AssetPriceDataSetIdBlockCompoundUniqueInput = {
    dataSetId: number
    block: number
  }

  export type AssetPriceCountOrderByAggregateInput = {
    id?: SortOrder
    dataSetId?: SortOrder
    createdAt?: SortOrder
    updatedAt?: SortOrder
    block?: SortOrder
    price?: SortOrder
    signature?: SortOrder
  }

  export type AssetPriceAvgOrderByAggregateInput = {
    id?: SortOrder
    dataSetId?: SortOrder
    block?: SortOrder
    price?: SortOrder
  }

  export type AssetPriceMaxOrderByAggregateInput = {
    id?: SortOrder
    dataSetId?: SortOrder
    createdAt?: SortOrder
    updatedAt?: SortOrder
    block?: SortOrder
    price?: SortOrder
    signature?: SortOrder
  }

  export type AssetPriceMinOrderByAggregateInput = {
    id?: SortOrder
    dataSetId?: SortOrder
    createdAt?: SortOrder
    updatedAt?: SortOrder
    block?: SortOrder
    price?: SortOrder
    signature?: SortOrder
  }

  export type AssetPriceSumOrderByAggregateInput = {
    id?: SortOrder
    dataSetId?: SortOrder
    block?: SortOrder
    price?: SortOrder
  }

  export type IntWithAggregatesFilter<$PrismaModel = never> = {
    equals?: number | IntFieldRefInput<$PrismaModel>
    in?: number[] | ListIntFieldRefInput<$PrismaModel>
    notIn?: number[] | ListIntFieldRefInput<$PrismaModel>
    lt?: number | IntFieldRefInput<$PrismaModel>
    lte?: number | IntFieldRefInput<$PrismaModel>
    gt?: number | IntFieldRefInput<$PrismaModel>
    gte?: number | IntFieldRefInput<$PrismaModel>
    not?: NestedIntWithAggregatesFilter<$PrismaModel> | number
    _count?: NestedIntFilter<$PrismaModel>
    _avg?: NestedFloatFilter<$PrismaModel>
    _sum?: NestedIntFilter<$PrismaModel>
    _min?: NestedIntFilter<$PrismaModel>
    _max?: NestedIntFilter<$PrismaModel>
  }

  export type DateTimeWithAggregatesFilter<$PrismaModel = never> = {
    equals?: Date | string | DateTimeFieldRefInput<$PrismaModel>
    in?: Date[] | string[] | ListDateTimeFieldRefInput<$PrismaModel>
    notIn?: Date[] | string[] | ListDateTimeFieldRefInput<$PrismaModel>
    lt?: Date | string | DateTimeFieldRefInput<$PrismaModel>
    lte?: Date | string | DateTimeFieldRefInput<$PrismaModel>
    gt?: Date | string | DateTimeFieldRefInput<$PrismaModel>
    gte?: Date | string | DateTimeFieldRefInput<$PrismaModel>
    not?: NestedDateTimeWithAggregatesFilter<$PrismaModel> | Date | string
    _count?: NestedIntFilter<$PrismaModel>
    _min?: NestedDateTimeFilter<$PrismaModel>
    _max?: NestedDateTimeFilter<$PrismaModel>
  }

  export type IntNullableWithAggregatesFilter<$PrismaModel = never> = {
    equals?: number | IntFieldRefInput<$PrismaModel> | null
    in?: number[] | ListIntFieldRefInput<$PrismaModel> | null
    notIn?: number[] | ListIntFieldRefInput<$PrismaModel> | null
    lt?: number | IntFieldRefInput<$PrismaModel>
    lte?: number | IntFieldRefInput<$PrismaModel>
    gt?: number | IntFieldRefInput<$PrismaModel>
    gte?: number | IntFieldRefInput<$PrismaModel>
    not?: NestedIntNullableWithAggregatesFilter<$PrismaModel> | number | null
    _count?: NestedIntNullableFilter<$PrismaModel>
    _avg?: NestedFloatNullableFilter<$PrismaModel>
    _sum?: NestedIntNullableFilter<$PrismaModel>
    _min?: NestedIntNullableFilter<$PrismaModel>
    _max?: NestedIntNullableFilter<$PrismaModel>
  }

  export type DecimalWithAggregatesFilter<$PrismaModel = never> = {
    equals?: Decimal | DecimalJsLike | number | string | DecimalFieldRefInput<$PrismaModel>
    in?: Decimal[] | DecimalJsLike[] | number[] | string[] | ListDecimalFieldRefInput<$PrismaModel>
    notIn?: Decimal[] | DecimalJsLike[] | number[] | string[] | ListDecimalFieldRefInput<$PrismaModel>
    lt?: Decimal | DecimalJsLike | number | string | DecimalFieldRefInput<$PrismaModel>
    lte?: Decimal | DecimalJsLike | number | string | DecimalFieldRefInput<$PrismaModel>
    gt?: Decimal | DecimalJsLike | number | string | DecimalFieldRefInput<$PrismaModel>
    gte?: Decimal | DecimalJsLike | number | string | DecimalFieldRefInput<$PrismaModel>
    not?: NestedDecimalWithAggregatesFilter<$PrismaModel> | Decimal | DecimalJsLike | number | string
    _count?: NestedIntFilter<$PrismaModel>
    _avg?: NestedDecimalFilter<$PrismaModel>
    _sum?: NestedDecimalFilter<$PrismaModel>
    _min?: NestedDecimalFilter<$PrismaModel>
    _max?: NestedDecimalFilter<$PrismaModel>
  }

  export type StringWithAggregatesFilter<$PrismaModel = never> = {
    equals?: string | StringFieldRefInput<$PrismaModel>
    in?: string[] | ListStringFieldRefInput<$PrismaModel>
    notIn?: string[] | ListStringFieldRefInput<$PrismaModel>
    lt?: string | StringFieldRefInput<$PrismaModel>
    lte?: string | StringFieldRefInput<$PrismaModel>
    gt?: string | StringFieldRefInput<$PrismaModel>
    gte?: string | StringFieldRefInput<$PrismaModel>
    contains?: string | StringFieldRefInput<$PrismaModel>
    startsWith?: string | StringFieldRefInput<$PrismaModel>
    endsWith?: string | StringFieldRefInput<$PrismaModel>
    mode?: QueryMode
    not?: NestedStringWithAggregatesFilter<$PrismaModel> | string
    _count?: NestedIntFilter<$PrismaModel>
    _min?: NestedStringFilter<$PrismaModel>
    _max?: NestedStringFilter<$PrismaModel>
  }

  export type AssetPriceListRelationFilter = {
    every?: AssetPriceWhereInput
    some?: AssetPriceWhereInput
    none?: AssetPriceWhereInput
  }

  export type AssetPriceOrderByRelationAggregateInput = {
    _count?: SortOrder
  }

  export type DataSetCountOrderByAggregateInput = {
    id?: SortOrder
    name?: SortOrder
  }

  export type DataSetAvgOrderByAggregateInput = {
    id?: SortOrder
  }

  export type DataSetMaxOrderByAggregateInput = {
    id?: SortOrder
    name?: SortOrder
  }

  export type DataSetMinOrderByAggregateInput = {
    id?: SortOrder
    name?: SortOrder
  }

  export type DataSetSumOrderByAggregateInput = {
    id?: SortOrder
  }

  export type StringNullableFilter<$PrismaModel = never> = {
    equals?: string | StringFieldRefInput<$PrismaModel> | null
    in?: string[] | ListStringFieldRefInput<$PrismaModel> | null
    notIn?: string[] | ListStringFieldRefInput<$PrismaModel> | null
    lt?: string | StringFieldRefInput<$PrismaModel>
    lte?: string | StringFieldRefInput<$PrismaModel>
    gt?: string | StringFieldRefInput<$PrismaModel>
    gte?: string | StringFieldRefInput<$PrismaModel>
    contains?: string | StringFieldRefInput<$PrismaModel>
    startsWith?: string | StringFieldRefInput<$PrismaModel>
    endsWith?: string | StringFieldRefInput<$PrismaModel>
    mode?: QueryMode
    not?: NestedStringNullableFilter<$PrismaModel> | string | null
  }

  export type SignerCountOrderByAggregateInput = {
    id?: SortOrder
    key?: SortOrder
    name?: SortOrder
  }

  export type SignerAvgOrderByAggregateInput = {
    id?: SortOrder
  }

  export type SignerMaxOrderByAggregateInput = {
    id?: SortOrder
    key?: SortOrder
    name?: SortOrder
  }

  export type SignerMinOrderByAggregateInput = {
    id?: SortOrder
    key?: SortOrder
    name?: SortOrder
  }

  export type SignerSumOrderByAggregateInput = {
    id?: SortOrder
  }

  export type StringNullableWithAggregatesFilter<$PrismaModel = never> = {
    equals?: string | StringFieldRefInput<$PrismaModel> | null
    in?: string[] | ListStringFieldRefInput<$PrismaModel> | null
    notIn?: string[] | ListStringFieldRefInput<$PrismaModel> | null
    lt?: string | StringFieldRefInput<$PrismaModel>
    lte?: string | StringFieldRefInput<$PrismaModel>
    gt?: string | StringFieldRefInput<$PrismaModel>
    gte?: string | StringFieldRefInput<$PrismaModel>
    contains?: string | StringFieldRefInput<$PrismaModel>
    startsWith?: string | StringFieldRefInput<$PrismaModel>
    endsWith?: string | StringFieldRefInput<$PrismaModel>
    mode?: QueryMode
    not?: NestedStringNullableWithAggregatesFilter<$PrismaModel> | string | null
    _count?: NestedIntNullableFilter<$PrismaModel>
    _min?: NestedStringNullableFilter<$PrismaModel>
    _max?: NestedStringNullableFilter<$PrismaModel>
  }

  export type AssetPriceRelationFilter = {
    is?: AssetPriceWhereInput
    isNot?: AssetPriceWhereInput
  }

  export type SignerRelationFilter = {
    is?: SignerWhereInput
    isNot?: SignerWhereInput
  }

  export type SignersOnAssetPriceSignerIdAssetPriceIdCompoundUniqueInput = {
    signerId: number
    assetPriceId: number
  }

  export type SignersOnAssetPriceCountOrderByAggregateInput = {
    signerId?: SortOrder
    assetPriceId?: SortOrder
  }

  export type SignersOnAssetPriceAvgOrderByAggregateInput = {
    signerId?: SortOrder
    assetPriceId?: SortOrder
  }

  export type SignersOnAssetPriceMaxOrderByAggregateInput = {
    signerId?: SortOrder
    assetPriceId?: SortOrder
  }

  export type SignersOnAssetPriceMinOrderByAggregateInput = {
    signerId?: SortOrder
    assetPriceId?: SortOrder
  }

  export type SignersOnAssetPriceSumOrderByAggregateInput = {
    signerId?: SortOrder
    assetPriceId?: SortOrder
  }

  export type DataSetCreateNestedOneWithoutAssetPriceInput = {
    create?: XOR<DataSetCreateWithoutAssetPriceInput, DataSetUncheckedCreateWithoutAssetPriceInput>
    connectOrCreate?: DataSetCreateOrConnectWithoutAssetPriceInput
    connect?: DataSetWhereUniqueInput
  }

  export type SignersOnAssetPriceCreateNestedManyWithoutAssetPriceInput = {
    create?: XOR<SignersOnAssetPriceCreateWithoutAssetPriceInput, SignersOnAssetPriceUncheckedCreateWithoutAssetPriceInput> | SignersOnAssetPriceCreateWithoutAssetPriceInput[] | SignersOnAssetPriceUncheckedCreateWithoutAssetPriceInput[]
    connectOrCreate?: SignersOnAssetPriceCreateOrConnectWithoutAssetPriceInput | SignersOnAssetPriceCreateOrConnectWithoutAssetPriceInput[]
    createMany?: SignersOnAssetPriceCreateManyAssetPriceInputEnvelope
    connect?: SignersOnAssetPriceWhereUniqueInput | SignersOnAssetPriceWhereUniqueInput[]
  }

  export type SignersOnAssetPriceUncheckedCreateNestedManyWithoutAssetPriceInput = {
    create?: XOR<SignersOnAssetPriceCreateWithoutAssetPriceInput, SignersOnAssetPriceUncheckedCreateWithoutAssetPriceInput> | SignersOnAssetPriceCreateWithoutAssetPriceInput[] | SignersOnAssetPriceUncheckedCreateWithoutAssetPriceInput[]
    connectOrCreate?: SignersOnAssetPriceCreateOrConnectWithoutAssetPriceInput | SignersOnAssetPriceCreateOrConnectWithoutAssetPriceInput[]
    createMany?: SignersOnAssetPriceCreateManyAssetPriceInputEnvelope
    connect?: SignersOnAssetPriceWhereUniqueInput | SignersOnAssetPriceWhereUniqueInput[]
  }

  export type DateTimeFieldUpdateOperationsInput = {
    set?: Date | string
  }

  export type NullableIntFieldUpdateOperationsInput = {
    set?: number | null
    increment?: number
    decrement?: number
    multiply?: number
    divide?: number
  }

  export type DecimalFieldUpdateOperationsInput = {
    set?: Decimal | DecimalJsLike | number | string
    increment?: Decimal | DecimalJsLike | number | string
    decrement?: Decimal | DecimalJsLike | number | string
    multiply?: Decimal | DecimalJsLike | number | string
    divide?: Decimal | DecimalJsLike | number | string
  }

  export type StringFieldUpdateOperationsInput = {
    set?: string
  }

  export type DataSetUpdateOneRequiredWithoutAssetPriceNestedInput = {
    create?: XOR<DataSetCreateWithoutAssetPriceInput, DataSetUncheckedCreateWithoutAssetPriceInput>
    connectOrCreate?: DataSetCreateOrConnectWithoutAssetPriceInput
    upsert?: DataSetUpsertWithoutAssetPriceInput
    connect?: DataSetWhereUniqueInput
    update?: XOR<XOR<DataSetUpdateToOneWithWhereWithoutAssetPriceInput, DataSetUpdateWithoutAssetPriceInput>, DataSetUncheckedUpdateWithoutAssetPriceInput>
  }

  export type SignersOnAssetPriceUpdateManyWithoutAssetPriceNestedInput = {
    create?: XOR<SignersOnAssetPriceCreateWithoutAssetPriceInput, SignersOnAssetPriceUncheckedCreateWithoutAssetPriceInput> | SignersOnAssetPriceCreateWithoutAssetPriceInput[] | SignersOnAssetPriceUncheckedCreateWithoutAssetPriceInput[]
    connectOrCreate?: SignersOnAssetPriceCreateOrConnectWithoutAssetPriceInput | SignersOnAssetPriceCreateOrConnectWithoutAssetPriceInput[]
    upsert?: SignersOnAssetPriceUpsertWithWhereUniqueWithoutAssetPriceInput | SignersOnAssetPriceUpsertWithWhereUniqueWithoutAssetPriceInput[]
    createMany?: SignersOnAssetPriceCreateManyAssetPriceInputEnvelope
    set?: SignersOnAssetPriceWhereUniqueInput | SignersOnAssetPriceWhereUniqueInput[]
    disconnect?: SignersOnAssetPriceWhereUniqueInput | SignersOnAssetPriceWhereUniqueInput[]
    delete?: SignersOnAssetPriceWhereUniqueInput | SignersOnAssetPriceWhereUniqueInput[]
    connect?: SignersOnAssetPriceWhereUniqueInput | SignersOnAssetPriceWhereUniqueInput[]
    update?: SignersOnAssetPriceUpdateWithWhereUniqueWithoutAssetPriceInput | SignersOnAssetPriceUpdateWithWhereUniqueWithoutAssetPriceInput[]
    updateMany?: SignersOnAssetPriceUpdateManyWithWhereWithoutAssetPriceInput | SignersOnAssetPriceUpdateManyWithWhereWithoutAssetPriceInput[]
    deleteMany?: SignersOnAssetPriceScalarWhereInput | SignersOnAssetPriceScalarWhereInput[]
  }

  export type IntFieldUpdateOperationsInput = {
    set?: number
    increment?: number
    decrement?: number
    multiply?: number
    divide?: number
  }

  export type SignersOnAssetPriceUncheckedUpdateManyWithoutAssetPriceNestedInput = {
    create?: XOR<SignersOnAssetPriceCreateWithoutAssetPriceInput, SignersOnAssetPriceUncheckedCreateWithoutAssetPriceInput> | SignersOnAssetPriceCreateWithoutAssetPriceInput[] | SignersOnAssetPriceUncheckedCreateWithoutAssetPriceInput[]
    connectOrCreate?: SignersOnAssetPriceCreateOrConnectWithoutAssetPriceInput | SignersOnAssetPriceCreateOrConnectWithoutAssetPriceInput[]
    upsert?: SignersOnAssetPriceUpsertWithWhereUniqueWithoutAssetPriceInput | SignersOnAssetPriceUpsertWithWhereUniqueWithoutAssetPriceInput[]
    createMany?: SignersOnAssetPriceCreateManyAssetPriceInputEnvelope
    set?: SignersOnAssetPriceWhereUniqueInput | SignersOnAssetPriceWhereUniqueInput[]
    disconnect?: SignersOnAssetPriceWhereUniqueInput | SignersOnAssetPriceWhereUniqueInput[]
    delete?: SignersOnAssetPriceWhereUniqueInput | SignersOnAssetPriceWhereUniqueInput[]
    connect?: SignersOnAssetPriceWhereUniqueInput | SignersOnAssetPriceWhereUniqueInput[]
    update?: SignersOnAssetPriceUpdateWithWhereUniqueWithoutAssetPriceInput | SignersOnAssetPriceUpdateWithWhereUniqueWithoutAssetPriceInput[]
    updateMany?: SignersOnAssetPriceUpdateManyWithWhereWithoutAssetPriceInput | SignersOnAssetPriceUpdateManyWithWhereWithoutAssetPriceInput[]
    deleteMany?: SignersOnAssetPriceScalarWhereInput | SignersOnAssetPriceScalarWhereInput[]
  }

  export type AssetPriceCreateNestedManyWithoutDatasetInput = {
    create?: XOR<AssetPriceCreateWithoutDatasetInput, AssetPriceUncheckedCreateWithoutDatasetInput> | AssetPriceCreateWithoutDatasetInput[] | AssetPriceUncheckedCreateWithoutDatasetInput[]
    connectOrCreate?: AssetPriceCreateOrConnectWithoutDatasetInput | AssetPriceCreateOrConnectWithoutDatasetInput[]
    createMany?: AssetPriceCreateManyDatasetInputEnvelope
    connect?: AssetPriceWhereUniqueInput | AssetPriceWhereUniqueInput[]
  }

  export type AssetPriceUncheckedCreateNestedManyWithoutDatasetInput = {
    create?: XOR<AssetPriceCreateWithoutDatasetInput, AssetPriceUncheckedCreateWithoutDatasetInput> | AssetPriceCreateWithoutDatasetInput[] | AssetPriceUncheckedCreateWithoutDatasetInput[]
    connectOrCreate?: AssetPriceCreateOrConnectWithoutDatasetInput | AssetPriceCreateOrConnectWithoutDatasetInput[]
    createMany?: AssetPriceCreateManyDatasetInputEnvelope
    connect?: AssetPriceWhereUniqueInput | AssetPriceWhereUniqueInput[]
  }

  export type AssetPriceUpdateManyWithoutDatasetNestedInput = {
    create?: XOR<AssetPriceCreateWithoutDatasetInput, AssetPriceUncheckedCreateWithoutDatasetInput> | AssetPriceCreateWithoutDatasetInput[] | AssetPriceUncheckedCreateWithoutDatasetInput[]
    connectOrCreate?: AssetPriceCreateOrConnectWithoutDatasetInput | AssetPriceCreateOrConnectWithoutDatasetInput[]
    upsert?: AssetPriceUpsertWithWhereUniqueWithoutDatasetInput | AssetPriceUpsertWithWhereUniqueWithoutDatasetInput[]
    createMany?: AssetPriceCreateManyDatasetInputEnvelope
    set?: AssetPriceWhereUniqueInput | AssetPriceWhereUniqueInput[]
    disconnect?: AssetPriceWhereUniqueInput | AssetPriceWhereUniqueInput[]
    delete?: AssetPriceWhereUniqueInput | AssetPriceWhereUniqueInput[]
    connect?: AssetPriceWhereUniqueInput | AssetPriceWhereUniqueInput[]
    update?: AssetPriceUpdateWithWhereUniqueWithoutDatasetInput | AssetPriceUpdateWithWhereUniqueWithoutDatasetInput[]
    updateMany?: AssetPriceUpdateManyWithWhereWithoutDatasetInput | AssetPriceUpdateManyWithWhereWithoutDatasetInput[]
    deleteMany?: AssetPriceScalarWhereInput | AssetPriceScalarWhereInput[]
  }

  export type AssetPriceUncheckedUpdateManyWithoutDatasetNestedInput = {
    create?: XOR<AssetPriceCreateWithoutDatasetInput, AssetPriceUncheckedCreateWithoutDatasetInput> | AssetPriceCreateWithoutDatasetInput[] | AssetPriceUncheckedCreateWithoutDatasetInput[]
    connectOrCreate?: AssetPriceCreateOrConnectWithoutDatasetInput | AssetPriceCreateOrConnectWithoutDatasetInput[]
    upsert?: AssetPriceUpsertWithWhereUniqueWithoutDatasetInput | AssetPriceUpsertWithWhereUniqueWithoutDatasetInput[]
    createMany?: AssetPriceCreateManyDatasetInputEnvelope
    set?: AssetPriceWhereUniqueInput | AssetPriceWhereUniqueInput[]
    disconnect?: AssetPriceWhereUniqueInput | AssetPriceWhereUniqueInput[]
    delete?: AssetPriceWhereUniqueInput | AssetPriceWhereUniqueInput[]
    connect?: AssetPriceWhereUniqueInput | AssetPriceWhereUniqueInput[]
    update?: AssetPriceUpdateWithWhereUniqueWithoutDatasetInput | AssetPriceUpdateWithWhereUniqueWithoutDatasetInput[]
    updateMany?: AssetPriceUpdateManyWithWhereWithoutDatasetInput | AssetPriceUpdateManyWithWhereWithoutDatasetInput[]
    deleteMany?: AssetPriceScalarWhereInput | AssetPriceScalarWhereInput[]
  }

  export type SignersOnAssetPriceCreateNestedManyWithoutSignerInput = {
    create?: XOR<SignersOnAssetPriceCreateWithoutSignerInput, SignersOnAssetPriceUncheckedCreateWithoutSignerInput> | SignersOnAssetPriceCreateWithoutSignerInput[] | SignersOnAssetPriceUncheckedCreateWithoutSignerInput[]
    connectOrCreate?: SignersOnAssetPriceCreateOrConnectWithoutSignerInput | SignersOnAssetPriceCreateOrConnectWithoutSignerInput[]
    createMany?: SignersOnAssetPriceCreateManySignerInputEnvelope
    connect?: SignersOnAssetPriceWhereUniqueInput | SignersOnAssetPriceWhereUniqueInput[]
  }

  export type SignersOnAssetPriceUncheckedCreateNestedManyWithoutSignerInput = {
    create?: XOR<SignersOnAssetPriceCreateWithoutSignerInput, SignersOnAssetPriceUncheckedCreateWithoutSignerInput> | SignersOnAssetPriceCreateWithoutSignerInput[] | SignersOnAssetPriceUncheckedCreateWithoutSignerInput[]
    connectOrCreate?: SignersOnAssetPriceCreateOrConnectWithoutSignerInput | SignersOnAssetPriceCreateOrConnectWithoutSignerInput[]
    createMany?: SignersOnAssetPriceCreateManySignerInputEnvelope
    connect?: SignersOnAssetPriceWhereUniqueInput | SignersOnAssetPriceWhereUniqueInput[]
  }

  export type NullableStringFieldUpdateOperationsInput = {
    set?: string | null
  }

  export type SignersOnAssetPriceUpdateManyWithoutSignerNestedInput = {
    create?: XOR<SignersOnAssetPriceCreateWithoutSignerInput, SignersOnAssetPriceUncheckedCreateWithoutSignerInput> | SignersOnAssetPriceCreateWithoutSignerInput[] | SignersOnAssetPriceUncheckedCreateWithoutSignerInput[]
    connectOrCreate?: SignersOnAssetPriceCreateOrConnectWithoutSignerInput | SignersOnAssetPriceCreateOrConnectWithoutSignerInput[]
    upsert?: SignersOnAssetPriceUpsertWithWhereUniqueWithoutSignerInput | SignersOnAssetPriceUpsertWithWhereUniqueWithoutSignerInput[]
    createMany?: SignersOnAssetPriceCreateManySignerInputEnvelope
    set?: SignersOnAssetPriceWhereUniqueInput | SignersOnAssetPriceWhereUniqueInput[]
    disconnect?: SignersOnAssetPriceWhereUniqueInput | SignersOnAssetPriceWhereUniqueInput[]
    delete?: SignersOnAssetPriceWhereUniqueInput | SignersOnAssetPriceWhereUniqueInput[]
    connect?: SignersOnAssetPriceWhereUniqueInput | SignersOnAssetPriceWhereUniqueInput[]
    update?: SignersOnAssetPriceUpdateWithWhereUniqueWithoutSignerInput | SignersOnAssetPriceUpdateWithWhereUniqueWithoutSignerInput[]
    updateMany?: SignersOnAssetPriceUpdateManyWithWhereWithoutSignerInput | SignersOnAssetPriceUpdateManyWithWhereWithoutSignerInput[]
    deleteMany?: SignersOnAssetPriceScalarWhereInput | SignersOnAssetPriceScalarWhereInput[]
  }

  export type SignersOnAssetPriceUncheckedUpdateManyWithoutSignerNestedInput = {
    create?: XOR<SignersOnAssetPriceCreateWithoutSignerInput, SignersOnAssetPriceUncheckedCreateWithoutSignerInput> | SignersOnAssetPriceCreateWithoutSignerInput[] | SignersOnAssetPriceUncheckedCreateWithoutSignerInput[]
    connectOrCreate?: SignersOnAssetPriceCreateOrConnectWithoutSignerInput | SignersOnAssetPriceCreateOrConnectWithoutSignerInput[]
    upsert?: SignersOnAssetPriceUpsertWithWhereUniqueWithoutSignerInput | SignersOnAssetPriceUpsertWithWhereUniqueWithoutSignerInput[]
    createMany?: SignersOnAssetPriceCreateManySignerInputEnvelope
    set?: SignersOnAssetPriceWhereUniqueInput | SignersOnAssetPriceWhereUniqueInput[]
    disconnect?: SignersOnAssetPriceWhereUniqueInput | SignersOnAssetPriceWhereUniqueInput[]
    delete?: SignersOnAssetPriceWhereUniqueInput | SignersOnAssetPriceWhereUniqueInput[]
    connect?: SignersOnAssetPriceWhereUniqueInput | SignersOnAssetPriceWhereUniqueInput[]
    update?: SignersOnAssetPriceUpdateWithWhereUniqueWithoutSignerInput | SignersOnAssetPriceUpdateWithWhereUniqueWithoutSignerInput[]
    updateMany?: SignersOnAssetPriceUpdateManyWithWhereWithoutSignerInput | SignersOnAssetPriceUpdateManyWithWhereWithoutSignerInput[]
    deleteMany?: SignersOnAssetPriceScalarWhereInput | SignersOnAssetPriceScalarWhereInput[]
  }

  export type AssetPriceCreateNestedOneWithoutSignersOnAssetPriceInput = {
    create?: XOR<AssetPriceCreateWithoutSignersOnAssetPriceInput, AssetPriceUncheckedCreateWithoutSignersOnAssetPriceInput>
    connectOrCreate?: AssetPriceCreateOrConnectWithoutSignersOnAssetPriceInput
    connect?: AssetPriceWhereUniqueInput
  }

  export type SignerCreateNestedOneWithoutSignersOnAssetPriceInput = {
    create?: XOR<SignerCreateWithoutSignersOnAssetPriceInput, SignerUncheckedCreateWithoutSignersOnAssetPriceInput>
    connectOrCreate?: SignerCreateOrConnectWithoutSignersOnAssetPriceInput
    connect?: SignerWhereUniqueInput
  }

  export type AssetPriceUpdateOneRequiredWithoutSignersOnAssetPriceNestedInput = {
    create?: XOR<AssetPriceCreateWithoutSignersOnAssetPriceInput, AssetPriceUncheckedCreateWithoutSignersOnAssetPriceInput>
    connectOrCreate?: AssetPriceCreateOrConnectWithoutSignersOnAssetPriceInput
    upsert?: AssetPriceUpsertWithoutSignersOnAssetPriceInput
    connect?: AssetPriceWhereUniqueInput
    update?: XOR<XOR<AssetPriceUpdateToOneWithWhereWithoutSignersOnAssetPriceInput, AssetPriceUpdateWithoutSignersOnAssetPriceInput>, AssetPriceUncheckedUpdateWithoutSignersOnAssetPriceInput>
  }

  export type SignerUpdateOneRequiredWithoutSignersOnAssetPriceNestedInput = {
    create?: XOR<SignerCreateWithoutSignersOnAssetPriceInput, SignerUncheckedCreateWithoutSignersOnAssetPriceInput>
    connectOrCreate?: SignerCreateOrConnectWithoutSignersOnAssetPriceInput
    upsert?: SignerUpsertWithoutSignersOnAssetPriceInput
    connect?: SignerWhereUniqueInput
    update?: XOR<XOR<SignerUpdateToOneWithWhereWithoutSignersOnAssetPriceInput, SignerUpdateWithoutSignersOnAssetPriceInput>, SignerUncheckedUpdateWithoutSignersOnAssetPriceInput>
  }

  export type NestedIntFilter<$PrismaModel = never> = {
    equals?: number | IntFieldRefInput<$PrismaModel>
    in?: number[] | ListIntFieldRefInput<$PrismaModel>
    notIn?: number[] | ListIntFieldRefInput<$PrismaModel>
    lt?: number | IntFieldRefInput<$PrismaModel>
    lte?: number | IntFieldRefInput<$PrismaModel>
    gt?: number | IntFieldRefInput<$PrismaModel>
    gte?: number | IntFieldRefInput<$PrismaModel>
    not?: NestedIntFilter<$PrismaModel> | number
  }

  export type NestedDateTimeFilter<$PrismaModel = never> = {
    equals?: Date | string | DateTimeFieldRefInput<$PrismaModel>
    in?: Date[] | string[] | ListDateTimeFieldRefInput<$PrismaModel>
    notIn?: Date[] | string[] | ListDateTimeFieldRefInput<$PrismaModel>
    lt?: Date | string | DateTimeFieldRefInput<$PrismaModel>
    lte?: Date | string | DateTimeFieldRefInput<$PrismaModel>
    gt?: Date | string | DateTimeFieldRefInput<$PrismaModel>
    gte?: Date | string | DateTimeFieldRefInput<$PrismaModel>
    not?: NestedDateTimeFilter<$PrismaModel> | Date | string
  }

  export type NestedIntNullableFilter<$PrismaModel = never> = {
    equals?: number | IntFieldRefInput<$PrismaModel> | null
    in?: number[] | ListIntFieldRefInput<$PrismaModel> | null
    notIn?: number[] | ListIntFieldRefInput<$PrismaModel> | null
    lt?: number | IntFieldRefInput<$PrismaModel>
    lte?: number | IntFieldRefInput<$PrismaModel>
    gt?: number | IntFieldRefInput<$PrismaModel>
    gte?: number | IntFieldRefInput<$PrismaModel>
    not?: NestedIntNullableFilter<$PrismaModel> | number | null
  }

  export type NestedDecimalFilter<$PrismaModel = never> = {
    equals?: Decimal | DecimalJsLike | number | string | DecimalFieldRefInput<$PrismaModel>
    in?: Decimal[] | DecimalJsLike[] | number[] | string[] | ListDecimalFieldRefInput<$PrismaModel>
    notIn?: Decimal[] | DecimalJsLike[] | number[] | string[] | ListDecimalFieldRefInput<$PrismaModel>
    lt?: Decimal | DecimalJsLike | number | string | DecimalFieldRefInput<$PrismaModel>
    lte?: Decimal | DecimalJsLike | number | string | DecimalFieldRefInput<$PrismaModel>
    gt?: Decimal | DecimalJsLike | number | string | DecimalFieldRefInput<$PrismaModel>
    gte?: Decimal | DecimalJsLike | number | string | DecimalFieldRefInput<$PrismaModel>
    not?: NestedDecimalFilter<$PrismaModel> | Decimal | DecimalJsLike | number | string
  }

  export type NestedStringFilter<$PrismaModel = never> = {
    equals?: string | StringFieldRefInput<$PrismaModel>
    in?: string[] | ListStringFieldRefInput<$PrismaModel>
    notIn?: string[] | ListStringFieldRefInput<$PrismaModel>
    lt?: string | StringFieldRefInput<$PrismaModel>
    lte?: string | StringFieldRefInput<$PrismaModel>
    gt?: string | StringFieldRefInput<$PrismaModel>
    gte?: string | StringFieldRefInput<$PrismaModel>
    contains?: string | StringFieldRefInput<$PrismaModel>
    startsWith?: string | StringFieldRefInput<$PrismaModel>
    endsWith?: string | StringFieldRefInput<$PrismaModel>
    not?: NestedStringFilter<$PrismaModel> | string
  }

  export type NestedIntWithAggregatesFilter<$PrismaModel = never> = {
    equals?: number | IntFieldRefInput<$PrismaModel>
    in?: number[] | ListIntFieldRefInput<$PrismaModel>
    notIn?: number[] | ListIntFieldRefInput<$PrismaModel>
    lt?: number | IntFieldRefInput<$PrismaModel>
    lte?: number | IntFieldRefInput<$PrismaModel>
    gt?: number | IntFieldRefInput<$PrismaModel>
    gte?: number | IntFieldRefInput<$PrismaModel>
    not?: NestedIntWithAggregatesFilter<$PrismaModel> | number
    _count?: NestedIntFilter<$PrismaModel>
    _avg?: NestedFloatFilter<$PrismaModel>
    _sum?: NestedIntFilter<$PrismaModel>
    _min?: NestedIntFilter<$PrismaModel>
    _max?: NestedIntFilter<$PrismaModel>
  }

  export type NestedFloatFilter<$PrismaModel = never> = {
    equals?: number | FloatFieldRefInput<$PrismaModel>
    in?: number[] | ListFloatFieldRefInput<$PrismaModel>
    notIn?: number[] | ListFloatFieldRefInput<$PrismaModel>
    lt?: number | FloatFieldRefInput<$PrismaModel>
    lte?: number | FloatFieldRefInput<$PrismaModel>
    gt?: number | FloatFieldRefInput<$PrismaModel>
    gte?: number | FloatFieldRefInput<$PrismaModel>
    not?: NestedFloatFilter<$PrismaModel> | number
  }

  export type NestedDateTimeWithAggregatesFilter<$PrismaModel = never> = {
    equals?: Date | string | DateTimeFieldRefInput<$PrismaModel>
    in?: Date[] | string[] | ListDateTimeFieldRefInput<$PrismaModel>
    notIn?: Date[] | string[] | ListDateTimeFieldRefInput<$PrismaModel>
    lt?: Date | string | DateTimeFieldRefInput<$PrismaModel>
    lte?: Date | string | DateTimeFieldRefInput<$PrismaModel>
    gt?: Date | string | DateTimeFieldRefInput<$PrismaModel>
    gte?: Date | string | DateTimeFieldRefInput<$PrismaModel>
    not?: NestedDateTimeWithAggregatesFilter<$PrismaModel> | Date | string
    _count?: NestedIntFilter<$PrismaModel>
    _min?: NestedDateTimeFilter<$PrismaModel>
    _max?: NestedDateTimeFilter<$PrismaModel>
  }

  export type NestedIntNullableWithAggregatesFilter<$PrismaModel = never> = {
    equals?: number | IntFieldRefInput<$PrismaModel> | null
    in?: number[] | ListIntFieldRefInput<$PrismaModel> | null
    notIn?: number[] | ListIntFieldRefInput<$PrismaModel> | null
    lt?: number | IntFieldRefInput<$PrismaModel>
    lte?: number | IntFieldRefInput<$PrismaModel>
    gt?: number | IntFieldRefInput<$PrismaModel>
    gte?: number | IntFieldRefInput<$PrismaModel>
    not?: NestedIntNullableWithAggregatesFilter<$PrismaModel> | number | null
    _count?: NestedIntNullableFilter<$PrismaModel>
    _avg?: NestedFloatNullableFilter<$PrismaModel>
    _sum?: NestedIntNullableFilter<$PrismaModel>
    _min?: NestedIntNullableFilter<$PrismaModel>
    _max?: NestedIntNullableFilter<$PrismaModel>
  }

  export type NestedFloatNullableFilter<$PrismaModel = never> = {
    equals?: number | FloatFieldRefInput<$PrismaModel> | null
    in?: number[] | ListFloatFieldRefInput<$PrismaModel> | null
    notIn?: number[] | ListFloatFieldRefInput<$PrismaModel> | null
    lt?: number | FloatFieldRefInput<$PrismaModel>
    lte?: number | FloatFieldRefInput<$PrismaModel>
    gt?: number | FloatFieldRefInput<$PrismaModel>
    gte?: number | FloatFieldRefInput<$PrismaModel>
    not?: NestedFloatNullableFilter<$PrismaModel> | number | null
  }

  export type NestedDecimalWithAggregatesFilter<$PrismaModel = never> = {
    equals?: Decimal | DecimalJsLike | number | string | DecimalFieldRefInput<$PrismaModel>
    in?: Decimal[] | DecimalJsLike[] | number[] | string[] | ListDecimalFieldRefInput<$PrismaModel>
    notIn?: Decimal[] | DecimalJsLike[] | number[] | string[] | ListDecimalFieldRefInput<$PrismaModel>
    lt?: Decimal | DecimalJsLike | number | string | DecimalFieldRefInput<$PrismaModel>
    lte?: Decimal | DecimalJsLike | number | string | DecimalFieldRefInput<$PrismaModel>
    gt?: Decimal | DecimalJsLike | number | string | DecimalFieldRefInput<$PrismaModel>
    gte?: Decimal | DecimalJsLike | number | string | DecimalFieldRefInput<$PrismaModel>
    not?: NestedDecimalWithAggregatesFilter<$PrismaModel> | Decimal | DecimalJsLike | number | string
    _count?: NestedIntFilter<$PrismaModel>
    _avg?: NestedDecimalFilter<$PrismaModel>
    _sum?: NestedDecimalFilter<$PrismaModel>
    _min?: NestedDecimalFilter<$PrismaModel>
    _max?: NestedDecimalFilter<$PrismaModel>
  }

  export type NestedStringWithAggregatesFilter<$PrismaModel = never> = {
    equals?: string | StringFieldRefInput<$PrismaModel>
    in?: string[] | ListStringFieldRefInput<$PrismaModel>
    notIn?: string[] | ListStringFieldRefInput<$PrismaModel>
    lt?: string | StringFieldRefInput<$PrismaModel>
    lte?: string | StringFieldRefInput<$PrismaModel>
    gt?: string | StringFieldRefInput<$PrismaModel>
    gte?: string | StringFieldRefInput<$PrismaModel>
    contains?: string | StringFieldRefInput<$PrismaModel>
    startsWith?: string | StringFieldRefInput<$PrismaModel>
    endsWith?: string | StringFieldRefInput<$PrismaModel>
    not?: NestedStringWithAggregatesFilter<$PrismaModel> | string
    _count?: NestedIntFilter<$PrismaModel>
    _min?: NestedStringFilter<$PrismaModel>
    _max?: NestedStringFilter<$PrismaModel>
  }

  export type NestedStringNullableFilter<$PrismaModel = never> = {
    equals?: string | StringFieldRefInput<$PrismaModel> | null
    in?: string[] | ListStringFieldRefInput<$PrismaModel> | null
    notIn?: string[] | ListStringFieldRefInput<$PrismaModel> | null
    lt?: string | StringFieldRefInput<$PrismaModel>
    lte?: string | StringFieldRefInput<$PrismaModel>
    gt?: string | StringFieldRefInput<$PrismaModel>
    gte?: string | StringFieldRefInput<$PrismaModel>
    contains?: string | StringFieldRefInput<$PrismaModel>
    startsWith?: string | StringFieldRefInput<$PrismaModel>
    endsWith?: string | StringFieldRefInput<$PrismaModel>
    not?: NestedStringNullableFilter<$PrismaModel> | string | null
  }

  export type NestedStringNullableWithAggregatesFilter<$PrismaModel = never> = {
    equals?: string | StringFieldRefInput<$PrismaModel> | null
    in?: string[] | ListStringFieldRefInput<$PrismaModel> | null
    notIn?: string[] | ListStringFieldRefInput<$PrismaModel> | null
    lt?: string | StringFieldRefInput<$PrismaModel>
    lte?: string | StringFieldRefInput<$PrismaModel>
    gt?: string | StringFieldRefInput<$PrismaModel>
    gte?: string | StringFieldRefInput<$PrismaModel>
    contains?: string | StringFieldRefInput<$PrismaModel>
    startsWith?: string | StringFieldRefInput<$PrismaModel>
    endsWith?: string | StringFieldRefInput<$PrismaModel>
    not?: NestedStringNullableWithAggregatesFilter<$PrismaModel> | string | null
    _count?: NestedIntNullableFilter<$PrismaModel>
    _min?: NestedStringNullableFilter<$PrismaModel>
    _max?: NestedStringNullableFilter<$PrismaModel>
  }

  export type DataSetCreateWithoutAssetPriceInput = {
    name: string
  }

  export type DataSetUncheckedCreateWithoutAssetPriceInput = {
    id?: number
    name: string
  }

  export type DataSetCreateOrConnectWithoutAssetPriceInput = {
    where: DataSetWhereUniqueInput
    create: XOR<DataSetCreateWithoutAssetPriceInput, DataSetUncheckedCreateWithoutAssetPriceInput>
  }

  export type SignersOnAssetPriceCreateWithoutAssetPriceInput = {
    signer: SignerCreateNestedOneWithoutSignersOnAssetPriceInput
  }

  export type SignersOnAssetPriceUncheckedCreateWithoutAssetPriceInput = {
    signerId: number
  }

  export type SignersOnAssetPriceCreateOrConnectWithoutAssetPriceInput = {
    where: SignersOnAssetPriceWhereUniqueInput
    create: XOR<SignersOnAssetPriceCreateWithoutAssetPriceInput, SignersOnAssetPriceUncheckedCreateWithoutAssetPriceInput>
  }

  export type SignersOnAssetPriceCreateManyAssetPriceInputEnvelope = {
    data: SignersOnAssetPriceCreateManyAssetPriceInput | SignersOnAssetPriceCreateManyAssetPriceInput[]
    skipDuplicates?: boolean
  }

  export type DataSetUpsertWithoutAssetPriceInput = {
    update: XOR<DataSetUpdateWithoutAssetPriceInput, DataSetUncheckedUpdateWithoutAssetPriceInput>
    create: XOR<DataSetCreateWithoutAssetPriceInput, DataSetUncheckedCreateWithoutAssetPriceInput>
    where?: DataSetWhereInput
  }

  export type DataSetUpdateToOneWithWhereWithoutAssetPriceInput = {
    where?: DataSetWhereInput
    data: XOR<DataSetUpdateWithoutAssetPriceInput, DataSetUncheckedUpdateWithoutAssetPriceInput>
  }

  export type DataSetUpdateWithoutAssetPriceInput = {
    name?: StringFieldUpdateOperationsInput | string
  }

  export type DataSetUncheckedUpdateWithoutAssetPriceInput = {
    id?: IntFieldUpdateOperationsInput | number
    name?: StringFieldUpdateOperationsInput | string
  }

  export type SignersOnAssetPriceUpsertWithWhereUniqueWithoutAssetPriceInput = {
    where: SignersOnAssetPriceWhereUniqueInput
    update: XOR<SignersOnAssetPriceUpdateWithoutAssetPriceInput, SignersOnAssetPriceUncheckedUpdateWithoutAssetPriceInput>
    create: XOR<SignersOnAssetPriceCreateWithoutAssetPriceInput, SignersOnAssetPriceUncheckedCreateWithoutAssetPriceInput>
  }

  export type SignersOnAssetPriceUpdateWithWhereUniqueWithoutAssetPriceInput = {
    where: SignersOnAssetPriceWhereUniqueInput
    data: XOR<SignersOnAssetPriceUpdateWithoutAssetPriceInput, SignersOnAssetPriceUncheckedUpdateWithoutAssetPriceInput>
  }

  export type SignersOnAssetPriceUpdateManyWithWhereWithoutAssetPriceInput = {
    where: SignersOnAssetPriceScalarWhereInput
    data: XOR<SignersOnAssetPriceUpdateManyMutationInput, SignersOnAssetPriceUncheckedUpdateManyWithoutAssetPriceInput>
  }

  export type SignersOnAssetPriceScalarWhereInput = {
    AND?: SignersOnAssetPriceScalarWhereInput | SignersOnAssetPriceScalarWhereInput[]
    OR?: SignersOnAssetPriceScalarWhereInput[]
    NOT?: SignersOnAssetPriceScalarWhereInput | SignersOnAssetPriceScalarWhereInput[]
    signerId?: IntFilter<"SignersOnAssetPrice"> | number
    assetPriceId?: IntFilter<"SignersOnAssetPrice"> | number
  }

  export type AssetPriceCreateWithoutDatasetInput = {
    createdAt?: Date | string
    updatedAt?: Date | string
    block?: number | null
    price: Decimal | DecimalJsLike | number | string
    signature: string
    signersOnAssetPrice?: SignersOnAssetPriceCreateNestedManyWithoutAssetPriceInput
  }

  export type AssetPriceUncheckedCreateWithoutDatasetInput = {
    id?: number
    createdAt?: Date | string
    updatedAt?: Date | string
    block?: number | null
    price: Decimal | DecimalJsLike | number | string
    signature: string
    signersOnAssetPrice?: SignersOnAssetPriceUncheckedCreateNestedManyWithoutAssetPriceInput
  }

  export type AssetPriceCreateOrConnectWithoutDatasetInput = {
    where: AssetPriceWhereUniqueInput
    create: XOR<AssetPriceCreateWithoutDatasetInput, AssetPriceUncheckedCreateWithoutDatasetInput>
  }

  export type AssetPriceCreateManyDatasetInputEnvelope = {
    data: AssetPriceCreateManyDatasetInput | AssetPriceCreateManyDatasetInput[]
    skipDuplicates?: boolean
  }

  export type AssetPriceUpsertWithWhereUniqueWithoutDatasetInput = {
    where: AssetPriceWhereUniqueInput
    update: XOR<AssetPriceUpdateWithoutDatasetInput, AssetPriceUncheckedUpdateWithoutDatasetInput>
    create: XOR<AssetPriceCreateWithoutDatasetInput, AssetPriceUncheckedCreateWithoutDatasetInput>
  }

  export type AssetPriceUpdateWithWhereUniqueWithoutDatasetInput = {
    where: AssetPriceWhereUniqueInput
    data: XOR<AssetPriceUpdateWithoutDatasetInput, AssetPriceUncheckedUpdateWithoutDatasetInput>
  }

  export type AssetPriceUpdateManyWithWhereWithoutDatasetInput = {
    where: AssetPriceScalarWhereInput
    data: XOR<AssetPriceUpdateManyMutationInput, AssetPriceUncheckedUpdateManyWithoutDatasetInput>
  }

  export type AssetPriceScalarWhereInput = {
    AND?: AssetPriceScalarWhereInput | AssetPriceScalarWhereInput[]
    OR?: AssetPriceScalarWhereInput[]
    NOT?: AssetPriceScalarWhereInput | AssetPriceScalarWhereInput[]
    id?: IntFilter<"AssetPrice"> | number
    dataSetId?: IntFilter<"AssetPrice"> | number
    createdAt?: DateTimeFilter<"AssetPrice"> | Date | string
    updatedAt?: DateTimeFilter<"AssetPrice"> | Date | string
    block?: IntNullableFilter<"AssetPrice"> | number | null
    price?: DecimalFilter<"AssetPrice"> | Decimal | DecimalJsLike | number | string
    signature?: StringFilter<"AssetPrice"> | string
  }

  export type SignersOnAssetPriceCreateWithoutSignerInput = {
    assetPrice: AssetPriceCreateNestedOneWithoutSignersOnAssetPriceInput
  }

  export type SignersOnAssetPriceUncheckedCreateWithoutSignerInput = {
    assetPriceId: number
  }

  export type SignersOnAssetPriceCreateOrConnectWithoutSignerInput = {
    where: SignersOnAssetPriceWhereUniqueInput
    create: XOR<SignersOnAssetPriceCreateWithoutSignerInput, SignersOnAssetPriceUncheckedCreateWithoutSignerInput>
  }

  export type SignersOnAssetPriceCreateManySignerInputEnvelope = {
    data: SignersOnAssetPriceCreateManySignerInput | SignersOnAssetPriceCreateManySignerInput[]
    skipDuplicates?: boolean
  }

  export type SignersOnAssetPriceUpsertWithWhereUniqueWithoutSignerInput = {
    where: SignersOnAssetPriceWhereUniqueInput
    update: XOR<SignersOnAssetPriceUpdateWithoutSignerInput, SignersOnAssetPriceUncheckedUpdateWithoutSignerInput>
    create: XOR<SignersOnAssetPriceCreateWithoutSignerInput, SignersOnAssetPriceUncheckedCreateWithoutSignerInput>
  }

  export type SignersOnAssetPriceUpdateWithWhereUniqueWithoutSignerInput = {
    where: SignersOnAssetPriceWhereUniqueInput
    data: XOR<SignersOnAssetPriceUpdateWithoutSignerInput, SignersOnAssetPriceUncheckedUpdateWithoutSignerInput>
  }

  export type SignersOnAssetPriceUpdateManyWithWhereWithoutSignerInput = {
    where: SignersOnAssetPriceScalarWhereInput
    data: XOR<SignersOnAssetPriceUpdateManyMutationInput, SignersOnAssetPriceUncheckedUpdateManyWithoutSignerInput>
  }

  export type AssetPriceCreateWithoutSignersOnAssetPriceInput = {
    createdAt?: Date | string
    updatedAt?: Date | string
    block?: number | null
    price: Decimal | DecimalJsLike | number | string
    signature: string
    dataset: DataSetCreateNestedOneWithoutAssetPriceInput
  }

  export type AssetPriceUncheckedCreateWithoutSignersOnAssetPriceInput = {
    id?: number
    dataSetId: number
    createdAt?: Date | string
    updatedAt?: Date | string
    block?: number | null
    price: Decimal | DecimalJsLike | number | string
    signature: string
  }

  export type AssetPriceCreateOrConnectWithoutSignersOnAssetPriceInput = {
    where: AssetPriceWhereUniqueInput
    create: XOR<AssetPriceCreateWithoutSignersOnAssetPriceInput, AssetPriceUncheckedCreateWithoutSignersOnAssetPriceInput>
  }

  export type SignerCreateWithoutSignersOnAssetPriceInput = {
    key: string
    name?: string | null
  }

  export type SignerUncheckedCreateWithoutSignersOnAssetPriceInput = {
    id?: number
    key: string
    name?: string | null
  }

  export type SignerCreateOrConnectWithoutSignersOnAssetPriceInput = {
    where: SignerWhereUniqueInput
    create: XOR<SignerCreateWithoutSignersOnAssetPriceInput, SignerUncheckedCreateWithoutSignersOnAssetPriceInput>
  }

  export type AssetPriceUpsertWithoutSignersOnAssetPriceInput = {
    update: XOR<AssetPriceUpdateWithoutSignersOnAssetPriceInput, AssetPriceUncheckedUpdateWithoutSignersOnAssetPriceInput>
    create: XOR<AssetPriceCreateWithoutSignersOnAssetPriceInput, AssetPriceUncheckedCreateWithoutSignersOnAssetPriceInput>
    where?: AssetPriceWhereInput
  }

  export type AssetPriceUpdateToOneWithWhereWithoutSignersOnAssetPriceInput = {
    where?: AssetPriceWhereInput
    data: XOR<AssetPriceUpdateWithoutSignersOnAssetPriceInput, AssetPriceUncheckedUpdateWithoutSignersOnAssetPriceInput>
  }

  export type AssetPriceUpdateWithoutSignersOnAssetPriceInput = {
    createdAt?: DateTimeFieldUpdateOperationsInput | Date | string
    updatedAt?: DateTimeFieldUpdateOperationsInput | Date | string
    block?: NullableIntFieldUpdateOperationsInput | number | null
    price?: DecimalFieldUpdateOperationsInput | Decimal | DecimalJsLike | number | string
    signature?: StringFieldUpdateOperationsInput | string
    dataset?: DataSetUpdateOneRequiredWithoutAssetPriceNestedInput
  }

  export type AssetPriceUncheckedUpdateWithoutSignersOnAssetPriceInput = {
    id?: IntFieldUpdateOperationsInput | number
    dataSetId?: IntFieldUpdateOperationsInput | number
    createdAt?: DateTimeFieldUpdateOperationsInput | Date | string
    updatedAt?: DateTimeFieldUpdateOperationsInput | Date | string
    block?: NullableIntFieldUpdateOperationsInput | number | null
    price?: DecimalFieldUpdateOperationsInput | Decimal | DecimalJsLike | number | string
    signature?: StringFieldUpdateOperationsInput | string
  }

  export type SignerUpsertWithoutSignersOnAssetPriceInput = {
    update: XOR<SignerUpdateWithoutSignersOnAssetPriceInput, SignerUncheckedUpdateWithoutSignersOnAssetPriceInput>
    create: XOR<SignerCreateWithoutSignersOnAssetPriceInput, SignerUncheckedCreateWithoutSignersOnAssetPriceInput>
    where?: SignerWhereInput
  }

  export type SignerUpdateToOneWithWhereWithoutSignersOnAssetPriceInput = {
    where?: SignerWhereInput
    data: XOR<SignerUpdateWithoutSignersOnAssetPriceInput, SignerUncheckedUpdateWithoutSignersOnAssetPriceInput>
  }

  export type SignerUpdateWithoutSignersOnAssetPriceInput = {
    key?: StringFieldUpdateOperationsInput | string
    name?: NullableStringFieldUpdateOperationsInput | string | null
  }

  export type SignerUncheckedUpdateWithoutSignersOnAssetPriceInput = {
    id?: IntFieldUpdateOperationsInput | number
    key?: StringFieldUpdateOperationsInput | string
    name?: NullableStringFieldUpdateOperationsInput | string | null
  }

  export type SignersOnAssetPriceCreateManyAssetPriceInput = {
    signerId: number
  }

  export type SignersOnAssetPriceUpdateWithoutAssetPriceInput = {
    signer?: SignerUpdateOneRequiredWithoutSignersOnAssetPriceNestedInput
  }

  export type SignersOnAssetPriceUncheckedUpdateWithoutAssetPriceInput = {
    signerId?: IntFieldUpdateOperationsInput | number
  }

  export type SignersOnAssetPriceUncheckedUpdateManyWithoutAssetPriceInput = {
    signerId?: IntFieldUpdateOperationsInput | number
  }

  export type AssetPriceCreateManyDatasetInput = {
    id?: number
    createdAt?: Date | string
    updatedAt?: Date | string
    block?: number | null
    price: Decimal | DecimalJsLike | number | string
    signature: string
  }

  export type AssetPriceUpdateWithoutDatasetInput = {
    createdAt?: DateTimeFieldUpdateOperationsInput | Date | string
    updatedAt?: DateTimeFieldUpdateOperationsInput | Date | string
    block?: NullableIntFieldUpdateOperationsInput | number | null
    price?: DecimalFieldUpdateOperationsInput | Decimal | DecimalJsLike | number | string
    signature?: StringFieldUpdateOperationsInput | string
    signersOnAssetPrice?: SignersOnAssetPriceUpdateManyWithoutAssetPriceNestedInput
  }

  export type AssetPriceUncheckedUpdateWithoutDatasetInput = {
    id?: IntFieldUpdateOperationsInput | number
    createdAt?: DateTimeFieldUpdateOperationsInput | Date | string
    updatedAt?: DateTimeFieldUpdateOperationsInput | Date | string
    block?: NullableIntFieldUpdateOperationsInput | number | null
    price?: DecimalFieldUpdateOperationsInput | Decimal | DecimalJsLike | number | string
    signature?: StringFieldUpdateOperationsInput | string
    signersOnAssetPrice?: SignersOnAssetPriceUncheckedUpdateManyWithoutAssetPriceNestedInput
  }

  export type AssetPriceUncheckedUpdateManyWithoutDatasetInput = {
    id?: IntFieldUpdateOperationsInput | number
    createdAt?: DateTimeFieldUpdateOperationsInput | Date | string
    updatedAt?: DateTimeFieldUpdateOperationsInput | Date | string
    block?: NullableIntFieldUpdateOperationsInput | number | null
    price?: DecimalFieldUpdateOperationsInput | Decimal | DecimalJsLike | number | string
    signature?: StringFieldUpdateOperationsInput | string
  }

  export type SignersOnAssetPriceCreateManySignerInput = {
    assetPriceId: number
  }

  export type SignersOnAssetPriceUpdateWithoutSignerInput = {
    assetPrice?: AssetPriceUpdateOneRequiredWithoutSignersOnAssetPriceNestedInput
  }

  export type SignersOnAssetPriceUncheckedUpdateWithoutSignerInput = {
    assetPriceId?: IntFieldUpdateOperationsInput | number
  }

  export type SignersOnAssetPriceUncheckedUpdateManyWithoutSignerInput = {
    assetPriceId?: IntFieldUpdateOperationsInput | number
  }



  /**
   * Aliases for legacy arg types
   */
    /**
     * @deprecated Use AssetPriceCountOutputTypeDefaultArgs instead
     */
    export type AssetPriceCountOutputTypeArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = AssetPriceCountOutputTypeDefaultArgs<ExtArgs>
    /**
     * @deprecated Use DataSetCountOutputTypeDefaultArgs instead
     */
    export type DataSetCountOutputTypeArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = DataSetCountOutputTypeDefaultArgs<ExtArgs>
    /**
     * @deprecated Use SignerCountOutputTypeDefaultArgs instead
     */
    export type SignerCountOutputTypeArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = SignerCountOutputTypeDefaultArgs<ExtArgs>
    /**
     * @deprecated Use AssetPriceDefaultArgs instead
     */
    export type AssetPriceArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = AssetPriceDefaultArgs<ExtArgs>
    /**
     * @deprecated Use DataSetDefaultArgs instead
     */
    export type DataSetArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = DataSetDefaultArgs<ExtArgs>
    /**
     * @deprecated Use SignerDefaultArgs instead
     */
    export type SignerArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = SignerDefaultArgs<ExtArgs>
    /**
     * @deprecated Use SignersOnAssetPriceDefaultArgs instead
     */
    export type SignersOnAssetPriceArgs<ExtArgs extends $Extensions.InternalArgs = $Extensions.DefaultArgs> = SignersOnAssetPriceDefaultArgs<ExtArgs>

  /**
   * Batch Payload for updateMany & deleteMany & createMany
   */

  export type BatchPayload = {
    count: number
  }

  /**
   * DMMF
   */
  export const dmmf: runtime.BaseDMMF
}