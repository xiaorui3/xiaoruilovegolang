# User ID Conversion and Error Handling Design

## Problem

`gin.Context.Param` returns the route parameter as a `string`, while
`model.User.SelectTestTableIdtoName` accepts an `int`. Passing the value
directly causes a compile-time type error. The model method also discards its
database error, so query failures are currently indistinguishable from an empty
name.

## Design

The controller will parse the route parameter with `strconv.Atoi` before
calling the model. If parsing fails, it will return the project's standard
error response with code 400 and stop processing.

The model method will return `(string, error)` and propagate the GORM query
error. The controller will translate a non-nil query error into the project's
standard error response with code 500. A successful query will continue to
return the name with the existing success response shape.

## Alternatives Considered

- Change the model method to accept `string`: rejected because database IDs are
  represented as `int` throughout the model and invalid input would reach the
  database layer.
- Ignore conversion and query errors: rejected because it can return a false
  success with an empty name.

## Verification

- A non-numeric route ID returns code 400 without querying the database.
- A valid numeric ID reaches the model as an integer.
- A database error returns code 500.
- A successful query returns the selected name.
- The project compiles and all Go tests pass.
