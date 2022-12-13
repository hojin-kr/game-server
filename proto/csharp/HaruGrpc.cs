// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: proto/haru.proto
// </auto-generated>
// Original file comments:
// Copyright 2015 gRPC authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
#pragma warning disable 0414, 1591
#region Designer generated code

using grpc = global::Grpc.Core;

namespace Haru {
  /// <summary>
  /// Service definition.
  /// </summary>
  public static partial class version1
  {
    static readonly string __ServiceName = "haru.version1";

    static readonly grpc::Marshaller<global::Haru.AccountRequest> __Marshaller_haru_AccountRequest = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::Haru.AccountRequest.Parser.ParseFrom);
    static readonly grpc::Marshaller<global::Haru.AccountReply> __Marshaller_haru_AccountReply = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::Haru.AccountReply.Parser.ParseFrom);
    static readonly grpc::Marshaller<global::Haru.ProfileRequest> __Marshaller_haru_ProfileRequest = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::Haru.ProfileRequest.Parser.ParseFrom);
    static readonly grpc::Marshaller<global::Haru.ProfileReply> __Marshaller_haru_ProfileReply = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::Haru.ProfileReply.Parser.ParseFrom);
    static readonly grpc::Marshaller<global::Haru.GameRequest> __Marshaller_haru_GameRequest = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::Haru.GameRequest.Parser.ParseFrom);
    static readonly grpc::Marshaller<global::Haru.GameReply> __Marshaller_haru_GameReply = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::Haru.GameReply.Parser.ParseFrom);
    static readonly grpc::Marshaller<global::Haru.FilterdGamesRequest> __Marshaller_haru_FilterdGamesRequest = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::Haru.FilterdGamesRequest.Parser.ParseFrom);
    static readonly grpc::Marshaller<global::Haru.FilterdGamesReply> __Marshaller_haru_FilterdGamesReply = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::Haru.FilterdGamesReply.Parser.ParseFrom);
    static readonly grpc::Marshaller<global::Haru.JoinRequest> __Marshaller_haru_JoinRequest = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::Haru.JoinRequest.Parser.ParseFrom);
    static readonly grpc::Marshaller<global::Haru.JoinReply> __Marshaller_haru_JoinReply = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::Haru.JoinReply.Parser.ParseFrom);
    static readonly grpc::Marshaller<global::Haru.ChatRequest> __Marshaller_haru_ChatRequest = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::Haru.ChatRequest.Parser.ParseFrom);
    static readonly grpc::Marshaller<global::Haru.ChatReply> __Marshaller_haru_ChatReply = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::Haru.ChatReply.Parser.ParseFrom);
    static readonly grpc::Marshaller<global::Haru.ChatMessageRequest> __Marshaller_haru_ChatMessageRequest = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::Haru.ChatMessageRequest.Parser.ParseFrom);

    static readonly grpc::Method<global::Haru.AccountRequest, global::Haru.AccountReply> __Method_CreateAccount = new grpc::Method<global::Haru.AccountRequest, global::Haru.AccountReply>(
        grpc::MethodType.Unary,
        __ServiceName,
        "CreateAccount",
        __Marshaller_haru_AccountRequest,
        __Marshaller_haru_AccountReply);

    static readonly grpc::Method<global::Haru.ProfileRequest, global::Haru.ProfileReply> __Method_GetProfile = new grpc::Method<global::Haru.ProfileRequest, global::Haru.ProfileReply>(
        grpc::MethodType.Unary,
        __ServiceName,
        "GetProfile",
        __Marshaller_haru_ProfileRequest,
        __Marshaller_haru_ProfileReply);

    static readonly grpc::Method<global::Haru.ProfileRequest, global::Haru.ProfileReply> __Method_UpdateProfile = new grpc::Method<global::Haru.ProfileRequest, global::Haru.ProfileReply>(
        grpc::MethodType.Unary,
        __ServiceName,
        "UpdateProfile",
        __Marshaller_haru_ProfileRequest,
        __Marshaller_haru_ProfileReply);

    static readonly grpc::Method<global::Haru.GameRequest, global::Haru.GameReply> __Method_CreateGame = new grpc::Method<global::Haru.GameRequest, global::Haru.GameReply>(
        grpc::MethodType.Unary,
        __ServiceName,
        "CreateGame",
        __Marshaller_haru_GameRequest,
        __Marshaller_haru_GameReply);

    static readonly grpc::Method<global::Haru.GameRequest, global::Haru.GameReply> __Method_UpdateGame = new grpc::Method<global::Haru.GameRequest, global::Haru.GameReply>(
        grpc::MethodType.Unary,
        __ServiceName,
        "UpdateGame",
        __Marshaller_haru_GameRequest,
        __Marshaller_haru_GameReply);

    static readonly grpc::Method<global::Haru.GameRequest, global::Haru.GameReply> __Method_GetGame = new grpc::Method<global::Haru.GameRequest, global::Haru.GameReply>(
        grpc::MethodType.Unary,
        __ServiceName,
        "GetGame",
        __Marshaller_haru_GameRequest,
        __Marshaller_haru_GameReply);

    static readonly grpc::Method<global::Haru.FilterdGamesRequest, global::Haru.FilterdGamesReply> __Method_GetFilterdGames = new grpc::Method<global::Haru.FilterdGamesRequest, global::Haru.FilterdGamesReply>(
        grpc::MethodType.Unary,
        __ServiceName,
        "GetFilterdGames",
        __Marshaller_haru_FilterdGamesRequest,
        __Marshaller_haru_FilterdGamesReply);

    static readonly grpc::Method<global::Haru.JoinRequest, global::Haru.JoinReply> __Method_Join = new grpc::Method<global::Haru.JoinRequest, global::Haru.JoinReply>(
        grpc::MethodType.Unary,
        __ServiceName,
        "Join",
        __Marshaller_haru_JoinRequest,
        __Marshaller_haru_JoinReply);

    static readonly grpc::Method<global::Haru.JoinRequest, global::Haru.JoinReply> __Method_GetMyJoins = new grpc::Method<global::Haru.JoinRequest, global::Haru.JoinReply>(
        grpc::MethodType.Unary,
        __ServiceName,
        "GetMyJoins",
        __Marshaller_haru_JoinRequest,
        __Marshaller_haru_JoinReply);

    static readonly grpc::Method<global::Haru.JoinRequest, global::Haru.JoinReply> __Method_GetGameJoins = new grpc::Method<global::Haru.JoinRequest, global::Haru.JoinReply>(
        grpc::MethodType.Unary,
        __ServiceName,
        "GetGameJoins",
        __Marshaller_haru_JoinRequest,
        __Marshaller_haru_JoinReply);

    static readonly grpc::Method<global::Haru.JoinRequest, global::Haru.JoinReply> __Method_UpdateJoin = new grpc::Method<global::Haru.JoinRequest, global::Haru.JoinReply>(
        grpc::MethodType.Unary,
        __ServiceName,
        "UpdateJoin",
        __Marshaller_haru_JoinRequest,
        __Marshaller_haru_JoinReply);

    static readonly grpc::Method<global::Haru.ChatRequest, global::Haru.ChatReply> __Method_GetChat = new grpc::Method<global::Haru.ChatRequest, global::Haru.ChatReply>(
        grpc::MethodType.Unary,
        __ServiceName,
        "GetChat",
        __Marshaller_haru_ChatRequest,
        __Marshaller_haru_ChatReply);

    static readonly grpc::Method<global::Haru.ChatMessageRequest, global::Haru.ChatReply> __Method_AddChatMessage = new grpc::Method<global::Haru.ChatMessageRequest, global::Haru.ChatReply>(
        grpc::MethodType.Unary,
        __ServiceName,
        "AddChatMessage",
        __Marshaller_haru_ChatMessageRequest,
        __Marshaller_haru_ChatReply);

    /// <summary>Service descriptor</summary>
    public static global::Google.Protobuf.Reflection.ServiceDescriptor Descriptor
    {
      get { return global::Haru.HaruReflection.Descriptor.Services[0]; }
    }

    /// <summary>Base class for server-side implementations of version1</summary>
    [grpc::BindServiceMethod(typeof(version1), "BindService")]
    public abstract partial class version1Base
    {
      public virtual global::System.Threading.Tasks.Task<global::Haru.AccountReply> CreateAccount(global::Haru.AccountRequest request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

      public virtual global::System.Threading.Tasks.Task<global::Haru.ProfileReply> GetProfile(global::Haru.ProfileRequest request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

      public virtual global::System.Threading.Tasks.Task<global::Haru.ProfileReply> UpdateProfile(global::Haru.ProfileRequest request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

      public virtual global::System.Threading.Tasks.Task<global::Haru.GameReply> CreateGame(global::Haru.GameRequest request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

      public virtual global::System.Threading.Tasks.Task<global::Haru.GameReply> UpdateGame(global::Haru.GameRequest request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

      public virtual global::System.Threading.Tasks.Task<global::Haru.GameReply> GetGame(global::Haru.GameRequest request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

      public virtual global::System.Threading.Tasks.Task<global::Haru.FilterdGamesReply> GetFilterdGames(global::Haru.FilterdGamesRequest request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

      public virtual global::System.Threading.Tasks.Task<global::Haru.JoinReply> Join(global::Haru.JoinRequest request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

      public virtual global::System.Threading.Tasks.Task<global::Haru.JoinReply> GetMyJoins(global::Haru.JoinRequest request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

      public virtual global::System.Threading.Tasks.Task<global::Haru.JoinReply> GetGameJoins(global::Haru.JoinRequest request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

      public virtual global::System.Threading.Tasks.Task<global::Haru.JoinReply> UpdateJoin(global::Haru.JoinRequest request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

      public virtual global::System.Threading.Tasks.Task<global::Haru.ChatReply> GetChat(global::Haru.ChatRequest request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

      public virtual global::System.Threading.Tasks.Task<global::Haru.ChatReply> AddChatMessage(global::Haru.ChatMessageRequest request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

    }

    /// <summary>Client for version1</summary>
    public partial class version1Client : grpc::ClientBase<version1Client>
    {
      /// <summary>Creates a new client for version1</summary>
      /// <param name="channel">The channel to use to make remote calls.</param>
      public version1Client(grpc::ChannelBase channel) : base(channel)
      {
      }
      /// <summary>Creates a new client for version1 that uses a custom <c>CallInvoker</c>.</summary>
      /// <param name="callInvoker">The callInvoker to use to make remote calls.</param>
      public version1Client(grpc::CallInvoker callInvoker) : base(callInvoker)
      {
      }
      /// <summary>Protected parameterless constructor to allow creation of test doubles.</summary>
      protected version1Client() : base()
      {
      }
      /// <summary>Protected constructor to allow creation of configured clients.</summary>
      /// <param name="configuration">The client configuration.</param>
      protected version1Client(ClientBaseConfiguration configuration) : base(configuration)
      {
      }

      public virtual global::Haru.AccountReply CreateAccount(global::Haru.AccountRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return CreateAccount(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual global::Haru.AccountReply CreateAccount(global::Haru.AccountRequest request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_CreateAccount, null, options, request);
      }
      public virtual grpc::AsyncUnaryCall<global::Haru.AccountReply> CreateAccountAsync(global::Haru.AccountRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return CreateAccountAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual grpc::AsyncUnaryCall<global::Haru.AccountReply> CreateAccountAsync(global::Haru.AccountRequest request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_CreateAccount, null, options, request);
      }
      public virtual global::Haru.ProfileReply GetProfile(global::Haru.ProfileRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return GetProfile(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual global::Haru.ProfileReply GetProfile(global::Haru.ProfileRequest request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_GetProfile, null, options, request);
      }
      public virtual grpc::AsyncUnaryCall<global::Haru.ProfileReply> GetProfileAsync(global::Haru.ProfileRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return GetProfileAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual grpc::AsyncUnaryCall<global::Haru.ProfileReply> GetProfileAsync(global::Haru.ProfileRequest request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_GetProfile, null, options, request);
      }
      public virtual global::Haru.ProfileReply UpdateProfile(global::Haru.ProfileRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return UpdateProfile(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual global::Haru.ProfileReply UpdateProfile(global::Haru.ProfileRequest request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_UpdateProfile, null, options, request);
      }
      public virtual grpc::AsyncUnaryCall<global::Haru.ProfileReply> UpdateProfileAsync(global::Haru.ProfileRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return UpdateProfileAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual grpc::AsyncUnaryCall<global::Haru.ProfileReply> UpdateProfileAsync(global::Haru.ProfileRequest request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_UpdateProfile, null, options, request);
      }
      public virtual global::Haru.GameReply CreateGame(global::Haru.GameRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return CreateGame(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual global::Haru.GameReply CreateGame(global::Haru.GameRequest request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_CreateGame, null, options, request);
      }
      public virtual grpc::AsyncUnaryCall<global::Haru.GameReply> CreateGameAsync(global::Haru.GameRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return CreateGameAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual grpc::AsyncUnaryCall<global::Haru.GameReply> CreateGameAsync(global::Haru.GameRequest request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_CreateGame, null, options, request);
      }
      public virtual global::Haru.GameReply UpdateGame(global::Haru.GameRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return UpdateGame(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual global::Haru.GameReply UpdateGame(global::Haru.GameRequest request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_UpdateGame, null, options, request);
      }
      public virtual grpc::AsyncUnaryCall<global::Haru.GameReply> UpdateGameAsync(global::Haru.GameRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return UpdateGameAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual grpc::AsyncUnaryCall<global::Haru.GameReply> UpdateGameAsync(global::Haru.GameRequest request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_UpdateGame, null, options, request);
      }
      public virtual global::Haru.GameReply GetGame(global::Haru.GameRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return GetGame(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual global::Haru.GameReply GetGame(global::Haru.GameRequest request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_GetGame, null, options, request);
      }
      public virtual grpc::AsyncUnaryCall<global::Haru.GameReply> GetGameAsync(global::Haru.GameRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return GetGameAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual grpc::AsyncUnaryCall<global::Haru.GameReply> GetGameAsync(global::Haru.GameRequest request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_GetGame, null, options, request);
      }
      public virtual global::Haru.FilterdGamesReply GetFilterdGames(global::Haru.FilterdGamesRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return GetFilterdGames(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual global::Haru.FilterdGamesReply GetFilterdGames(global::Haru.FilterdGamesRequest request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_GetFilterdGames, null, options, request);
      }
      public virtual grpc::AsyncUnaryCall<global::Haru.FilterdGamesReply> GetFilterdGamesAsync(global::Haru.FilterdGamesRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return GetFilterdGamesAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual grpc::AsyncUnaryCall<global::Haru.FilterdGamesReply> GetFilterdGamesAsync(global::Haru.FilterdGamesRequest request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_GetFilterdGames, null, options, request);
      }
      public virtual global::Haru.JoinReply Join(global::Haru.JoinRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return Join(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual global::Haru.JoinReply Join(global::Haru.JoinRequest request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_Join, null, options, request);
      }
      public virtual grpc::AsyncUnaryCall<global::Haru.JoinReply> JoinAsync(global::Haru.JoinRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return JoinAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual grpc::AsyncUnaryCall<global::Haru.JoinReply> JoinAsync(global::Haru.JoinRequest request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_Join, null, options, request);
      }
      public virtual global::Haru.JoinReply GetMyJoins(global::Haru.JoinRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return GetMyJoins(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual global::Haru.JoinReply GetMyJoins(global::Haru.JoinRequest request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_GetMyJoins, null, options, request);
      }
      public virtual grpc::AsyncUnaryCall<global::Haru.JoinReply> GetMyJoinsAsync(global::Haru.JoinRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return GetMyJoinsAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual grpc::AsyncUnaryCall<global::Haru.JoinReply> GetMyJoinsAsync(global::Haru.JoinRequest request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_GetMyJoins, null, options, request);
      }
      public virtual global::Haru.JoinReply GetGameJoins(global::Haru.JoinRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return GetGameJoins(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual global::Haru.JoinReply GetGameJoins(global::Haru.JoinRequest request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_GetGameJoins, null, options, request);
      }
      public virtual grpc::AsyncUnaryCall<global::Haru.JoinReply> GetGameJoinsAsync(global::Haru.JoinRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return GetGameJoinsAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual grpc::AsyncUnaryCall<global::Haru.JoinReply> GetGameJoinsAsync(global::Haru.JoinRequest request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_GetGameJoins, null, options, request);
      }
      public virtual global::Haru.JoinReply UpdateJoin(global::Haru.JoinRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return UpdateJoin(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual global::Haru.JoinReply UpdateJoin(global::Haru.JoinRequest request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_UpdateJoin, null, options, request);
      }
      public virtual grpc::AsyncUnaryCall<global::Haru.JoinReply> UpdateJoinAsync(global::Haru.JoinRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return UpdateJoinAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual grpc::AsyncUnaryCall<global::Haru.JoinReply> UpdateJoinAsync(global::Haru.JoinRequest request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_UpdateJoin, null, options, request);
      }
      public virtual global::Haru.ChatReply GetChat(global::Haru.ChatRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return GetChat(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual global::Haru.ChatReply GetChat(global::Haru.ChatRequest request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_GetChat, null, options, request);
      }
      public virtual grpc::AsyncUnaryCall<global::Haru.ChatReply> GetChatAsync(global::Haru.ChatRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return GetChatAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual grpc::AsyncUnaryCall<global::Haru.ChatReply> GetChatAsync(global::Haru.ChatRequest request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_GetChat, null, options, request);
      }
      public virtual global::Haru.ChatReply AddChatMessage(global::Haru.ChatMessageRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return AddChatMessage(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual global::Haru.ChatReply AddChatMessage(global::Haru.ChatMessageRequest request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_AddChatMessage, null, options, request);
      }
      public virtual grpc::AsyncUnaryCall<global::Haru.ChatReply> AddChatMessageAsync(global::Haru.ChatMessageRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return AddChatMessageAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual grpc::AsyncUnaryCall<global::Haru.ChatReply> AddChatMessageAsync(global::Haru.ChatMessageRequest request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_AddChatMessage, null, options, request);
      }
      /// <summary>Creates a new instance of client from given <c>ClientBaseConfiguration</c>.</summary>
      protected override version1Client NewInstance(ClientBaseConfiguration configuration)
      {
        return new version1Client(configuration);
      }
    }

    /// <summary>Creates service definition that can be registered with a server</summary>
    /// <param name="serviceImpl">An object implementing the server-side handling logic.</param>
    public static grpc::ServerServiceDefinition BindService(version1Base serviceImpl)
    {
      return grpc::ServerServiceDefinition.CreateBuilder()
          .AddMethod(__Method_CreateAccount, serviceImpl.CreateAccount)
          .AddMethod(__Method_GetProfile, serviceImpl.GetProfile)
          .AddMethod(__Method_UpdateProfile, serviceImpl.UpdateProfile)
          .AddMethod(__Method_CreateGame, serviceImpl.CreateGame)
          .AddMethod(__Method_UpdateGame, serviceImpl.UpdateGame)
          .AddMethod(__Method_GetGame, serviceImpl.GetGame)
          .AddMethod(__Method_GetFilterdGames, serviceImpl.GetFilterdGames)
          .AddMethod(__Method_Join, serviceImpl.Join)
          .AddMethod(__Method_GetMyJoins, serviceImpl.GetMyJoins)
          .AddMethod(__Method_GetGameJoins, serviceImpl.GetGameJoins)
          .AddMethod(__Method_UpdateJoin, serviceImpl.UpdateJoin)
          .AddMethod(__Method_GetChat, serviceImpl.GetChat)
          .AddMethod(__Method_AddChatMessage, serviceImpl.AddChatMessage).Build();
    }

    /// <summary>Register service method with a service binder with or without implementation. Useful when customizing the  service binding logic.
    /// Note: this method is part of an experimental API that can change or be removed without any prior notice.</summary>
    /// <param name="serviceBinder">Service methods will be bound by calling <c>AddMethod</c> on this object.</param>
    /// <param name="serviceImpl">An object implementing the server-side handling logic.</param>
    public static void BindService(grpc::ServiceBinderBase serviceBinder, version1Base serviceImpl)
    {
      serviceBinder.AddMethod(__Method_CreateAccount, serviceImpl == null ? null : new grpc::UnaryServerMethod<global::Haru.AccountRequest, global::Haru.AccountReply>(serviceImpl.CreateAccount));
      serviceBinder.AddMethod(__Method_GetProfile, serviceImpl == null ? null : new grpc::UnaryServerMethod<global::Haru.ProfileRequest, global::Haru.ProfileReply>(serviceImpl.GetProfile));
      serviceBinder.AddMethod(__Method_UpdateProfile, serviceImpl == null ? null : new grpc::UnaryServerMethod<global::Haru.ProfileRequest, global::Haru.ProfileReply>(serviceImpl.UpdateProfile));
      serviceBinder.AddMethod(__Method_CreateGame, serviceImpl == null ? null : new grpc::UnaryServerMethod<global::Haru.GameRequest, global::Haru.GameReply>(serviceImpl.CreateGame));
      serviceBinder.AddMethod(__Method_UpdateGame, serviceImpl == null ? null : new grpc::UnaryServerMethod<global::Haru.GameRequest, global::Haru.GameReply>(serviceImpl.UpdateGame));
      serviceBinder.AddMethod(__Method_GetGame, serviceImpl == null ? null : new grpc::UnaryServerMethod<global::Haru.GameRequest, global::Haru.GameReply>(serviceImpl.GetGame));
      serviceBinder.AddMethod(__Method_GetFilterdGames, serviceImpl == null ? null : new grpc::UnaryServerMethod<global::Haru.FilterdGamesRequest, global::Haru.FilterdGamesReply>(serviceImpl.GetFilterdGames));
      serviceBinder.AddMethod(__Method_Join, serviceImpl == null ? null : new grpc::UnaryServerMethod<global::Haru.JoinRequest, global::Haru.JoinReply>(serviceImpl.Join));
      serviceBinder.AddMethod(__Method_GetMyJoins, serviceImpl == null ? null : new grpc::UnaryServerMethod<global::Haru.JoinRequest, global::Haru.JoinReply>(serviceImpl.GetMyJoins));
      serviceBinder.AddMethod(__Method_GetGameJoins, serviceImpl == null ? null : new grpc::UnaryServerMethod<global::Haru.JoinRequest, global::Haru.JoinReply>(serviceImpl.GetGameJoins));
      serviceBinder.AddMethod(__Method_UpdateJoin, serviceImpl == null ? null : new grpc::UnaryServerMethod<global::Haru.JoinRequest, global::Haru.JoinReply>(serviceImpl.UpdateJoin));
      serviceBinder.AddMethod(__Method_GetChat, serviceImpl == null ? null : new grpc::UnaryServerMethod<global::Haru.ChatRequest, global::Haru.ChatReply>(serviceImpl.GetChat));
      serviceBinder.AddMethod(__Method_AddChatMessage, serviceImpl == null ? null : new grpc::UnaryServerMethod<global::Haru.ChatMessageRequest, global::Haru.ChatReply>(serviceImpl.AddChatMessage));
    }

  }
}
#endregion
