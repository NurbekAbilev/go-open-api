# Project for learning purposes

- [x] Change to MUX
- [ ] Create Auth/Middelware
- [ ] Database migration    
- [x] Host swagger UI
- [x] Use ENV file
- [ ] Change Database Connection to PGX [link](https://github.com/jackc/pgx)
- [ ] Add indexes to firstname, lastname on employees
- [ ] Add updated_at created_at columns
- [ ] Use make Target [link](https://www.gnu.org/software/make/manual/html_node/Rule-Introduction.html)
- [ ] Update created_at updated_at columns after table change
- [ ] Add support for pagination
- [ ] Add error handler
- [ ] Write Integrations/HTTP tests ([] make them seperate from unit)

## Useful links Open API
- https://github.com/deepmap/oapi-codegen
- https://github.com/getkin/kin-openapi
- https://github.com/swaggest/swgui

## Migrations
- Create migration 

```sh
make create_migration name=${migration_name}
```

## List of questions

1. Test folder vs name_test
2. Use .env vs use os env
3. How to create make with parameters
