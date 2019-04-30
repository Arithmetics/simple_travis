psql -U postgres -d travis_ci_test -c "CREATE TABLE fun (cool int, verycool int);"

psql -U postgres -d travis_ci_test -c "INSERT INTO fun (cool, verycool) VALUES (55, 888);"

