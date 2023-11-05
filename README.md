# Nøtes

> The plain text note experience

## TO DO

- Doc comments
- Unit test
- pagination
- Refactor pipiline (serialize | validate | State management | serialize)
- Footer with tags
- Login
- better style
- animate card transition
- theme management + custom css
- support JSON  for endpoints
- Packaging
- Migration management
- import export from other note app
- Delete confirm
- Updated / Deleted successfully notification
- Move package app to app/server 
- Move openAPi doc comments from main to server
- Move static to server/static
- Add config for db Path
- Group rooting by ressource

## FIX 
 
 - error page, logging and swagger doc

## Generate OpenAPI Specs

```bash
swag init --output app/static/openapi/ --outputTypes "json,yaml"
```

## CLI 

notes-server                                    show help
notes-server init                               initialize the server
notes-server run                                start the server
notes-server db                                 show help for db migration
notes-server db migrate                         migrate to latest schema version.
notes-server db upgrade                         migrate to the next available schema version
notes-server db rollback                        migrate to the previous schema version
notes-server users                              show help for users administration
notes-server users list                         list users
notes-server users create NAME PWD              create a new user
notes-server users update ID NEW_NAME NEW_PWD   update user name and pasword
notes-server users delete ID                    delete user but keep data. add --purge to delete data.


## Acknownledge

CSS Reset from [Andy-Bell](https://andy-bell.co.uk/a-more-modern-css-reset/)

JS Masonry adapted from ["A Lightweight Masonry Solution" by  Ana Tudor](https://css-tricks.com/a-lightweight-masonry-solution/) 
