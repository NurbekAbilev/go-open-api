# Project for learning purposes

## General
- [x] Use ENV file
- [x] Change to MUX
- [ ] Add error handler
- [ ] Add support for pagination
- [ ] Use make Target [link](https://www.gnu.org/software/make/manual/html_node/Rule-Introduction.html)

## Database
- [x] Database migration    
- [x] Add indexes to firstname, lastname on employees
- [ ] Change Database Connection to PGX [link](https://github.com/jackc/pgx)
- [ ] Update updated_at created_at columns

## Open Api
- [x] Host swagger UI

## Testing
- [ ] Write Integrations/HTTP tests ([] make them seperate from unit)
- [ ] Create Auth/Middelware

## Useful links Open API
- https://github.com/deepmap/oapi-codegen
- https://github.com/getkin/kin-openapi
- https://github.com/swaggest/swgui

## Migrations
Create migration 

```sh
make create_migration name=${migration_name}
```

