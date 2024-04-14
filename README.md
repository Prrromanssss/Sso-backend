# Sso-backend

```
syntax = "proto3";

package auth;

option go_package = "prrromanssss.sso.v1;ssov1";

service Auth {
    rpc Register (RegisterRequest) returns (RegisterResponse);
    rpc Login(LoginRequest) returns (LoginResponse);
    rpc IsAdmin(IsAdminRequest) returns (IsAdminResponse);
}

message RegisterRequest {
    string email = 1;  // Email of the user to register.
    string password = 2;  // Password of ther user to register.
}

message RegisterResponse {
    int64 user_id = 1;  // User ID of the registered user.
}

message LoginRequest {
    string email = 1;  // Email of the user to login.
    string password = 2;  // Password of ther user to login.
    int32 app_id = 3;  // ID of the app to login.
}

message LoginResponse {
    string token = 1;  // ID token of the logged user.
}


message IsAdminRequest {
    int64 user_id = 1;  // User ID to validate.
}

message IsAdminResponse {
    bool is_admin = 1;  // Indicates whether the user is admin.
}
```
