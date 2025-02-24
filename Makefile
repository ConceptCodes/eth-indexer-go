POSTGRES_USER=postgres
POSTGRES_DB=eth_indexer
DUMP_FILE=eth_indexer_dump.sql
HOST=localhost
PORT=5432

dump-db:
	@echo "Dumping database..."
	@pg_dump -h $(HOST) -p $(PORT) -U $(POSTGRES_USER) $(POSTGRES_DB) > $(DUMP_FILE)
	@echo "Database dumped to $(DUMP_FILE)"

clear-db:
	@echo "Clearing all rows from all tables..."
	@psql -h $(HOST) -p $(PORT) -U $(POSTGRES_USER) -d $(POSTGRES_DB) -c "DO \$$\$ DECLARE r RECORD; BEGIN FOR r IN (SELECT tablename FROM pg_tables WHERE schemaname = current_schema()) LOOP EXECUTE 'TRUNCATE TABLE ' || quote_ident(r.tablename) || ' CASCADE'; END LOOP; END \$$\$;"
	@echo "All rows cleared from all tables."