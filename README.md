# Sso-backend

Proto file

- Service to authorization

```
service Auth {
    rpc Register (RegisterRequest) returns (RegisterResponse);
    rpc Login(LoginRequest) returns (LoginResponse);
    rpc IsAdmin(IsAdminRequest) returns (IsAdminResponse);
}
```
- Messages to register new user
```
message RegisterRequest {
    string email = 1;  // Email of the user to register.
    string password = 2;  // Password of ther user to register.
}

message RegisterResponse {
    int64 user_id = 1;  // User ID of the registered user.
}
```
- Messages to login user
```
message LoginRequest {
    string email = 1;  // Email of the user to login.
    string password = 2;  // Password of ther user to login.
    int32 app_id = 3;  // ID of the app to login.
}

message LoginResponse {
    string token = 1;  // ID token of the logged user.
}
```
- Messages to check if user is admin
```
message IsAdminRequest {
    int64 user_id = 1;  // User ID to validate.
}

message IsAdminResponse {
    bool is_admin = 1;  // Indicates whether the user is admin.
}
```
