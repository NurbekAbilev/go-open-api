# Project for learning purposes

- [ ] Run migrations from code, so I can run migrations in tests epta
- [ ] Setup TestMain function for setup/teardown
- [ ] Add testify lib

## General
- [x] Use ENV file
- [x] Change to MUX
- [ ] Add error handler
- [ ] Add support for pagination
- [x] Use make Target [link](https://www.gnu.org/software/make/manual/html_node/Rule-Introduction.html)
- [ ] Hash passwords

## Database
- [x] Database migration    
- [x] Add indexes to firstname, lastname on employees
- [ ] Change Database Connection to PGX [link](https://github.com/jackc/pgx)
- [ ] Update updated_at created_at columns
- [ ] Run tests un parallel using t.Parallel() (each test in its own schema) [link](https://medium.com/kongkow-it-medan/parallel-database-integration-test-on-go-application-8706b150ee2e)

## Open Api
- [x] Host swagger UI

## Testing
- [ ] Write Integrations/HTTP tests (make them seperate from unit)
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

## Drop all test schemas
```SQL
DO $$ 
DECLARE 
    s_name text;
BEGIN
    FOR s_name IN (SELECT schema_name FROM information_schema.schemata WHERE schema_name LIKE 'test\_%') LOOP
        -- Generate and execute the DROP SCHEMA statement for each matching schema.
        EXECUTE 'DROP SCHEMA IF EXISTS ' || s_name || ' CASCADE';
        RAISE NOTICE 'Dropped schema: %', s_name;
    END LOOP;
END $$;
```

