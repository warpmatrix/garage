= Rust

== Error Handling

- `ok`: Converts from `Result<T, E>` to `Option<T>`
- `unwrap`: Returns the contained `Ok` value, consuming the `self` value; Otherwise, panic and abort the entire program
- `expect`: Returns the contained `Ok` value, consuming the `self` value; Otherwise, panic with message and abort the entire program
