# Task Manager API – Unit Testing and CI Report

## Objective
To ensure correctness and stability of the Task Manager API by writing **unit tests** for key components using the `testify` and `mockery` libraries, and integrating tests into the **GitHub Actions CI pipeline**.

---

## Test Coverage Summary

| Package                               | Coverage       |
|--------------------------------------|----------------|
| Delivery/controllers                 | ✅ 100.0%       |
| Usecases                              | ✅ ~83.3%%       |
| Infrastructure                        | ✅ ~85.0%       |
| Infrastructure/middleware             | ✅ 100.0%       |


Run command:
```bash
go test ./... -cover

## Testing Strategy

###  Frameworks Used:
- **testify** – assertions and test suite  
- **mockery** – auto-generate mocks from interfaces  
- Native Go testing package (`testing`)  

### Mocking Strategy
- Mocked routers when testing controllers to isolate controller logic from routing concerns.  
- Mocked repository interfaces when testing usecases.  
- Used `mockery` to auto-generate mocks into the `/mocks` folder.  
- This ensured isolation of units and reproducibility.  

### What Was Tested

| Layer           | Tested       | Description                                         |
|-----------------|--------------|-----------------------------------------------------|
| Controllers     | ✅ Unit tested | Business logic routing + edge cases, routers mocked |
| Usecases       | ✅ Mock tested | Core logic tested with mocked Repositories          |
| Middleware (JWT) | ✅ Tested     | Token verification, auth logic                       |
| Infrastructure  | ✅ Partial    | Password handling, helper services                   |
| Routers         | ❌ No direct tests | Covered indirectly via mocked routers in controller tests |
| Repositories    | ✅ Indirectly | Mocked in usecase tests; direct DB logic skipped    |

## Notes

- Routers are not directly tested as they mainly declare route bindings, but they are exercised via controller tests where routers are mocked.  
- MongoDB repository interfaces were mocked to isolate usecases from the database.  
- No direct integration or E2E tests were added as per the current scope.  
- Clean Architecture and dependency injection helped improve testability.  
