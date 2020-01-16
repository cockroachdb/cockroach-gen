// Code generated by help.awk. DO NOT EDIT.
// GENERATED FILE DO NOT EDIT

package parser

var helpMessages = map[string]HelpMessageBody{
	//line sql.y: 1134
	`ALTER`: {
		//line sql.y: 1135
		Category: hGroup,
		//line sql.y: 1136
		Text: `ALTER TABLE, ALTER INDEX, ALTER VIEW, ALTER SEQUENCE, ALTER DATABASE, ALTER USER
`,
	},
	//line sql.y: 1151
	`ALTER TABLE`: {
		ShortDescription: `change the definition of a table`,
		//line sql.y: 1152
		Category: hDDL,
		//line sql.y: 1153
		Text: `
ALTER TABLE [IF EXISTS] <tablename> <command> [, ...]

Commands:
  ALTER TABLE ... ADD [COLUMN] [IF NOT EXISTS] <colname> <type> [<qualifiers...>]
  ALTER TABLE ... ADD <constraint>
  ALTER TABLE ... DROP [COLUMN] [IF EXISTS] <colname> [RESTRICT | CASCADE]
  ALTER TABLE ... DROP CONSTRAINT [IF EXISTS] <constraintname> [RESTRICT | CASCADE]
  ALTER TABLE ... ALTER [COLUMN] <colname> {SET DEFAULT <expr> | DROP DEFAULT}
  ALTER TABLE ... ALTER [COLUMN] <colname> DROP NOT NULL
  ALTER TABLE ... ALTER [COLUMN] <colname> DROP STORED
  ALTER TABLE ... ALTER [COLUMN] <colname> [SET DATA] TYPE <type> [COLLATE <collation>]
  ALTER TABLE ... ALTER PRIMARY KEY USING INDEX <name>
  ALTER TABLE ... RENAME TO <newname>
  ALTER TABLE ... RENAME [COLUMN] <colname> TO <newname>
  ALTER TABLE ... VALIDATE CONSTRAINT <constraintname>
  ALTER TABLE ... SPLIT AT <selectclause> [WITH EXPIRATION <expr>]
  ALTER TABLE ... UNSPLIT AT <selectclause>
  ALTER TABLE ... UNSPLIT ALL
  ALTER TABLE ... SCATTER [ FROM ( <exprs...> ) TO ( <exprs...> ) ]
  ALTER TABLE ... INJECT STATISTICS ...  (experimental)
  ALTER TABLE ... PARTITION BY RANGE ( <name...> ) ( <rangespec> )
  ALTER TABLE ... PARTITION BY LIST ( <name...> ) ( <listspec> )
  ALTER TABLE ... PARTITION BY NOTHING
  ALTER TABLE ... CONFIGURE ZONE <zoneconfig>

Column qualifiers:
  [CONSTRAINT <constraintname>] {NULL | NOT NULL | UNIQUE | PRIMARY KEY | CHECK (<expr>) | DEFAULT <expr>}
  FAMILY <familyname>, CREATE [IF NOT EXISTS] FAMILY [<familyname>]
  REFERENCES <tablename> [( <colnames...> )]
  COLLATE <collationname>

Zone configurations:
  DISCARD
  USING <var> = <expr> [, ...]
  USING <var> = COPY FROM PARENT [, ...]
  { TO | = } <expr>

`,
		//line sql.y: 1191
		SeeAlso: `WEBDOCS/alter-table.html
`,
	},
	//line sql.y: 1205
	`ALTER PARTITION`: {
		ShortDescription: `apply zone configurations to a partition`,
		//line sql.y: 1206
		Category: hDDL,
		//line sql.y: 1207
		Text: `
ALTER PARTITION <name> <command>

Commands:
  -- Alter a single partition which exists on any of a table's indexes.
  ALTER PARTITION <partition> OF TABLE <tablename> CONFIGURE ZONE <zoneconfig>

  -- Alter a partition of a specific index.
  ALTER PARTITION <partition> OF INDEX <tablename>@<indexname> CONFIGURE ZONE <zoneconfig>

  -- Alter all partitions with the same name across a table's indexes.
  ALTER PARTITION <partition> OF INDEX <tablename>@* CONFIGURE ZONE <zoneconfig>

Zone configurations:
  DISCARD
  USING <var> = <expr> [, ...]
  USING <var> = COPY FROM PARENT [, ...]
  { TO | = } <expr>

`,
		//line sql.y: 1226
		SeeAlso: `WEBDOCS/configure-zone.html
`,
	},
	//line sql.y: 1231
	`ALTER VIEW`: {
		ShortDescription: `change the definition of a view`,
		//line sql.y: 1232
		Category: hDDL,
		//line sql.y: 1233
		Text: `
ALTER VIEW [IF EXISTS] <name> RENAME TO <newname>
`,
		//line sql.y: 1235
		SeeAlso: `WEBDOCS/alter-view.html
`,
	},
	//line sql.y: 1242
	`ALTER SEQUENCE`: {
		ShortDescription: `change the definition of a sequence`,
		//line sql.y: 1243
		Category: hDDL,
		//line sql.y: 1244
		Text: `
ALTER SEQUENCE [IF EXISTS] <name>
  [INCREMENT <increment>]
  [MINVALUE <minvalue> | NO MINVALUE]
  [MAXVALUE <maxvalue> | NO MAXVALUE]
  [START <start>]
  [[NO] CYCLE]
ALTER SEQUENCE [IF EXISTS] <name> RENAME TO <newname>
`,
	},
	//line sql.y: 1267
	`ALTER USER`: {
		ShortDescription: `change user properties`,
		//line sql.y: 1268
		Category: hPriv,
		//line sql.y: 1269
		Text: `
ALTER USER [IF EXISTS] <name> WITH PASSWORD <password>
`,
		//line sql.y: 1271
		SeeAlso: `CREATE USER
`,
	},
	//line sql.y: 1276
	`ALTER DATABASE`: {
		ShortDescription: `change the definition of a database`,
		//line sql.y: 1277
		Category: hDDL,
		//line sql.y: 1278
		Text: `
ALTER DATABASE <name> RENAME TO <newname>
`,
		//line sql.y: 1280
		SeeAlso: `WEBDOCS/alter-database.html
`,
	},
	//line sql.y: 1288
	`ALTER RANGE`: {
		ShortDescription: `change the parameters of a range`,
		//line sql.y: 1289
		Category: hDDL,
		//line sql.y: 1290
		Text: `
ALTER RANGE <zonename> <command>

Commands:
  ALTER RANGE ... CONFIGURE ZONE <zoneconfig>

Zone configurations:
  DISCARD
  USING <var> = <expr> [, ...]
  USING <var> = COPY FROM PARENT [, ...]
  { TO | = } <expr>

`,
		//line sql.y: 1302
		SeeAlso: `ALTER TABLE
`,
	},
	//line sql.y: 1307
	`ALTER INDEX`: {
		ShortDescription: `change the definition of an index`,
		//line sql.y: 1308
		Category: hDDL,
		//line sql.y: 1309
		Text: `
ALTER INDEX [IF EXISTS] <idxname> <command>

Commands:
  ALTER INDEX ... RENAME TO <newname>
  ALTER INDEX ... SPLIT AT <selectclause> [WITH EXPIRATION <expr>]
  ALTER INDEX ... UNSPLIT AT <selectclause>
  ALTER INDEX ... UNSPLIT ALL
  ALTER INDEX ... SCATTER [ FROM ( <exprs...> ) TO ( <exprs...> ) ]

Zone configurations:
  DISCARD
  USING <var> = <expr> [, ...]
  USING <var> = COPY FROM PARENT [, ...]
  { TO | = } <expr>

`,
		//line sql.y: 1325
		SeeAlso: `WEBDOCS/alter-index.html
`,
	},
	//line sql.y: 1826
	`BACKUP`: {
		ShortDescription: `back up data to external storage`,
		//line sql.y: 1827
		Category: hCCL,
		//line sql.y: 1828
		Text: `
BACKUP <targets...> TO <location...>
       [ AS OF SYSTEM TIME <expr> ]
       [ INCREMENTAL FROM <location...> ]
       [ WITH <option> [= <value>] [, ...] ]

Targets:
   TABLE <pattern> [, ...]
   DATABASE <databasename> [, ...]

Location:
   "[scheme]://[host]/[path to backup]?[parameters]"

Options:
   INTO_DB
   SKIP_MISSING_FOREIGN_KEYS

`,
		//line sql.y: 1845
		SeeAlso: `RESTORE, WEBDOCS/backup.html
`,
	},
	//line sql.y: 1857
	`RESTORE`: {
		ShortDescription: `restore data from external storage`,
		//line sql.y: 1858
		Category: hCCL,
		//line sql.y: 1859
		Text: `
RESTORE <targets...> FROM <location...>
        [ AS OF SYSTEM TIME <expr> ]
        [ WITH <option> [= <value>] [, ...] ]

Targets:
   TABLE <pattern> [, ...]
   DATABASE <databasename> [, ...]

Locations:
   "[scheme]://[host]/[path to backup]?[parameters]"

Options:
   INTO_DB
   SKIP_MISSING_FOREIGN_KEYS

`,
		//line sql.y: 1875
		SeeAlso: `BACKUP, WEBDOCS/restore.html
`,
	},
	//line sql.y: 1909
	`IMPORT`: {
		ShortDescription: `load data from file in a distributed manner`,
		//line sql.y: 1910
		Category: hCCL,
		//line sql.y: 1911
		Text: `
-- Import both schema and table data:
IMPORT [ TABLE <tablename> FROM ]
       <format> <datafile>
       [ WITH <option> [= <value>] [, ...] ]

-- Import using specific schema, use only table data from external file:
IMPORT TABLE <tablename>
       { ( <elements> ) | CREATE USING <schemafile> }
       <format>
       DATA ( <datafile> [, ...] )
       [ WITH <option> [= <value>] [, ...] ]

Formats:
   CSV
   DELIMITED
   MYSQLDUMP
   PGCOPY
   PGDUMP

Options:
   distributed = '...'
   sstsize = '...'
   temp = '...'
   delimiter = '...'      [CSV, PGCOPY-specific]
   nullif = '...'         [CSV, PGCOPY-specific]
   comment = '...'        [CSV-specific]

`,
		//line sql.y: 1939
		SeeAlso: `CREATE TABLE
`,
	},
	//line sql.y: 1983
	`EXPORT`: {
		ShortDescription: `export data to file in a distributed manner`,
		//line sql.y: 1984
		Category: hCCL,
		//line sql.y: 1985
		Text: `
EXPORT INTO <format> <datafile> [WITH <option> [= value] [,...]] FROM <query>

Formats:
   CSV

Options:
   delimiter = '...'   [CSV-specific]

`,
		//line sql.y: 1994
		SeeAlso: `SELECT
`,
	},
	//line sql.y: 2088
	`CANCEL`: {
		//line sql.y: 2089
		Category: hGroup,
		//line sql.y: 2090
		Text: `CANCEL JOBS, CANCEL QUERIES, CANCEL SESSIONS
`,
	},
	//line sql.y: 2097
	`CANCEL JOBS`: {
		ShortDescription: `cancel background jobs`,
		//line sql.y: 2098
		Category: hMisc,
		//line sql.y: 2099
		Text: `
CANCEL JOBS <selectclause>
CANCEL JOB <jobid>
`,
		//line sql.y: 2102
		SeeAlso: `SHOW JOBS, PAUSE JOBS, RESUME JOBS
`,
	},
	//line sql.y: 2120
	`CANCEL QUERIES`: {
		ShortDescription: `cancel running queries`,
		//line sql.y: 2121
		Category: hMisc,
		//line sql.y: 2122
		Text: `
CANCEL QUERIES [IF EXISTS] <selectclause>
CANCEL QUERY [IF EXISTS] <expr>
`,
		//line sql.y: 2125
		SeeAlso: `SHOW QUERIES
`,
	},
	//line sql.y: 2156
	`CANCEL SESSIONS`: {
		ShortDescription: `cancel open sessions`,
		//line sql.y: 2157
		Category: hMisc,
		//line sql.y: 2158
		Text: `
CANCEL SESSIONS [IF EXISTS] <selectclause>
CANCEL SESSION [IF EXISTS] <sessionid>
`,
		//line sql.y: 2161
		SeeAlso: `SHOW SESSIONS
`,
	},
	//line sql.y: 2231
	`CREATE`: {
		//line sql.y: 2232
		Category: hGroup,
		//line sql.y: 2233
		Text: `
CREATE DATABASE, CREATE TABLE, CREATE INDEX, CREATE TABLE AS,
CREATE USER, CREATE VIEW, CREATE SEQUENCE, CREATE STATISTICS,
CREATE ROLE
`,
	},
	//line sql.y: 2314
	`CREATE STATISTICS`: {
		ShortDescription: `create a new table statistic`,
		//line sql.y: 2315
		Category: hMisc,
		//line sql.y: 2316
		Text: `
CREATE STATISTICS <statisticname>
  [ON <colname> [, ...]]
  FROM <tablename> [AS OF SYSTEM TIME <expr>]
`,
	},
	//line sql.y: 2459
	`DELETE`: {
		ShortDescription: `delete rows from a table`,
		//line sql.y: 2460
		Category: hDML,
		//line sql.y: 2461
		Text: `DELETE FROM <tablename> [WHERE <expr>]
              [ORDER BY <exprs...>]
              [LIMIT <expr>]
              [RETURNING <exprs...>]
`,
		//line sql.y: 2465
		SeeAlso: `WEBDOCS/delete.html
`,
	},
	//line sql.y: 2485
	`DISCARD`: {
		ShortDescription: `reset the session to its initial state`,
		//line sql.y: 2486
		Category: hCfg,
		//line sql.y: 2487
		Text: `DISCARD ALL
`,
	},
	//line sql.y: 2499
	`DROP`: {
		//line sql.y: 2500
		Category: hGroup,
		//line sql.y: 2501
		Text: `
DROP DATABASE, DROP INDEX, DROP TABLE, DROP VIEW, DROP SEQUENCE,
DROP USER, DROP ROLE
`,
	},
	//line sql.y: 2518
	`DROP VIEW`: {
		ShortDescription: `remove a view`,
		//line sql.y: 2519
		Category: hDDL,
		//line sql.y: 2520
		Text: `DROP VIEW [IF EXISTS] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 2521
		SeeAlso: `WEBDOCS/drop-index.html
`,
	},
	//line sql.y: 2533
	`DROP SEQUENCE`: {
		ShortDescription: `remove a sequence`,
		//line sql.y: 2534
		Category: hDDL,
		//line sql.y: 2535
		Text: `DROP SEQUENCE [IF EXISTS] <sequenceName> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 2536
		SeeAlso: `DROP
`,
	},
	//line sql.y: 2548
	`DROP TABLE`: {
		ShortDescription: `remove a table`,
		//line sql.y: 2549
		Category: hDDL,
		//line sql.y: 2550
		Text: `DROP TABLE [IF EXISTS] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 2551
		SeeAlso: `WEBDOCS/drop-table.html
`,
	},
	//line sql.y: 2563
	`DROP INDEX`: {
		ShortDescription: `remove an index`,
		//line sql.y: 2564
		Category: hDDL,
		//line sql.y: 2565
		Text: `DROP INDEX [IF EXISTS] <idxname> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 2566
		SeeAlso: `WEBDOCS/drop-index.html
`,
	},
	//line sql.y: 2586
	`DROP DATABASE`: {
		ShortDescription: `remove a database`,
		//line sql.y: 2587
		Category: hDDL,
		//line sql.y: 2588
		Text: `DROP DATABASE [IF EXISTS] <databasename> [CASCADE | RESTRICT]
`,
		//line sql.y: 2589
		SeeAlso: `WEBDOCS/drop-database.html
`,
	},
	//line sql.y: 2609
	`DROP USER`: {
		ShortDescription: `remove a user`,
		//line sql.y: 2610
		Category: hPriv,
		//line sql.y: 2611
		Text: `DROP USER [IF EXISTS] <user> [, ...]
`,
		//line sql.y: 2612
		SeeAlso: `CREATE USER, SHOW USERS
`,
	},
	//line sql.y: 2624
	`DROP ROLE`: {
		ShortDescription: `remove a role`,
		//line sql.y: 2625
		Category: hPriv,
		//line sql.y: 2626
		Text: `DROP ROLE [IF EXISTS] <role> [, ...]
`,
		//line sql.y: 2627
		SeeAlso: `CREATE ROLE, SHOW ROLES
`,
	},
	//line sql.y: 2651
	`EXPLAIN`: {
		ShortDescription: `show the logical plan of a query`,
		//line sql.y: 2652
		Category: hMisc,
		//line sql.y: 2653
		Text: `
EXPLAIN <statement>
EXPLAIN ([PLAN ,] <planoptions...> ) <statement>
EXPLAIN [ANALYZE] (DISTSQL) <statement>
EXPLAIN ANALYZE [(DISTSQL)] <statement>

Explainable statements:
    SELECT, CREATE, DROP, ALTER, INSERT, UPSERT, UPDATE, DELETE,
    SHOW, EXPLAIN

Plan options:
    TYPES, VERBOSE, OPT

`,
		//line sql.y: 2666
		SeeAlso: `WEBDOCS/explain.html
`,
	},
	//line sql.y: 2749
	`PREPARE`: {
		ShortDescription: `prepare a statement for later execution`,
		//line sql.y: 2750
		Category: hMisc,
		//line sql.y: 2751
		Text: `PREPARE <name> [ ( <types...> ) ] AS <query>
`,
		//line sql.y: 2752
		SeeAlso: `EXECUTE, DEALLOCATE, DISCARD
`,
	},
	//line sql.y: 2783
	`EXECUTE`: {
		ShortDescription: `execute a statement prepared previously`,
		//line sql.y: 2784
		Category: hMisc,
		//line sql.y: 2785
		Text: `EXECUTE <name> [ ( <exprs...> ) ]
`,
		//line sql.y: 2786
		SeeAlso: `PREPARE, DEALLOCATE, DISCARD
`,
	},
	//line sql.y: 2816
	`DEALLOCATE`: {
		ShortDescription: `remove a prepared statement`,
		//line sql.y: 2817
		Category: hMisc,
		//line sql.y: 2818
		Text: `DEALLOCATE [PREPARE] { <name> | ALL }
`,
		//line sql.y: 2819
		SeeAlso: `PREPARE, EXECUTE, DISCARD
`,
	},
	//line sql.y: 2839
	`GRANT`: {
		ShortDescription: `define access privileges and role memberships`,
		//line sql.y: 2840
		Category: hPriv,
		//line sql.y: 2841
		Text: `
Grant privileges:
  GRANT {ALL | <privileges...> } ON <targets...> TO <grantees...>
Grant role membership (CCL only):
  GRANT <roles...> TO <grantees...> [WITH ADMIN OPTION]

Privileges:
  CREATE, DROP, GRANT, SELECT, INSERT, DELETE, UPDATE

Targets:
  DATABASE <databasename> [, ...]
  [TABLE] [<databasename> .] { <tablename> | * } [, ...]

`,
		//line sql.y: 2854
		SeeAlso: `REVOKE, WEBDOCS/grant.html
`,
	},
	//line sql.y: 2870
	`REVOKE`: {
		ShortDescription: `remove access privileges and role memberships`,
		//line sql.y: 2871
		Category: hPriv,
		//line sql.y: 2872
		Text: `
Revoke privileges:
  REVOKE {ALL | <privileges...> } ON <targets...> FROM <grantees...>
Revoke role membership (CCL only):
  REVOKE [ADMIN OPTION FOR] <roles...> FROM <grantees...>

Privileges:
  CREATE, DROP, GRANT, SELECT, INSERT, DELETE, UPDATE

Targets:
  DATABASE <databasename> [, <databasename>]...
  [TABLE] [<databasename> .] { <tablename> | * } [, ...]

`,
		//line sql.y: 2885
		SeeAlso: `GRANT, WEBDOCS/revoke.html
`,
	},
	//line sql.y: 2939
	`RESET`: {
		ShortDescription: `reset a session variable to its default value`,
		//line sql.y: 2940
		Category: hCfg,
		//line sql.y: 2941
		Text: `RESET [SESSION] <var>
`,
		//line sql.y: 2942
		SeeAlso: `RESET CLUSTER SETTING, WEBDOCS/set-vars.html
`,
	},
	//line sql.y: 2954
	`RESET CLUSTER SETTING`: {
		ShortDescription: `reset a cluster setting to its default value`,
		//line sql.y: 2955
		Category: hCfg,
		//line sql.y: 2956
		Text: `RESET CLUSTER SETTING <var>
`,
		//line sql.y: 2957
		SeeAlso: `SET CLUSTER SETTING, RESET
`,
	},
	//line sql.y: 2966
	`USE`: {
		ShortDescription: `set the current database`,
		//line sql.y: 2967
		Category: hCfg,
		//line sql.y: 2968
		Text: `USE <dbname>

"USE <dbname>" is an alias for "SET [SESSION] database = <dbname>".
`,
		//line sql.y: 2971
		SeeAlso: `SET SESSION, WEBDOCS/set-vars.html
`,
	},
	//line sql.y: 2992
	`SCRUB`: {
		ShortDescription: `run checks against databases or tables`,
		//line sql.y: 2993
		Category: hExperimental,
		//line sql.y: 2994
		Text: `
EXPERIMENTAL SCRUB TABLE <table> ...
EXPERIMENTAL SCRUB DATABASE <database>

The various checks that ca be run with SCRUB includes:
  - Physical table data (encoding)
  - Secondary index integrity
  - Constraint integrity (NOT NULL, CHECK, FOREIGN KEY, UNIQUE)
`,
		//line sql.y: 3002
		SeeAlso: `SCRUB TABLE, SCRUB DATABASE
`,
	},
	//line sql.y: 3008
	`SCRUB DATABASE`: {
		ShortDescription: `run scrub checks on a database`,
		//line sql.y: 3009
		Category: hExperimental,
		//line sql.y: 3010
		Text: `
EXPERIMENTAL SCRUB DATABASE <database>
                            [AS OF SYSTEM TIME <expr>]

All scrub checks will be run on the database. This includes:
  - Physical table data (encoding)
  - Secondary index integrity
  - Constraint integrity (NOT NULL, CHECK, FOREIGN KEY, UNIQUE)
`,
		//line sql.y: 3018
		SeeAlso: `SCRUB TABLE, SCRUB
`,
	},
	//line sql.y: 3026
	`SCRUB TABLE`: {
		ShortDescription: `run scrub checks on a table`,
		//line sql.y: 3027
		Category: hExperimental,
		//line sql.y: 3028
		Text: `
SCRUB TABLE <tablename>
            [AS OF SYSTEM TIME <expr>]
            [WITH OPTIONS <option> [, ...]]

Options:
  EXPERIMENTAL SCRUB TABLE ... WITH OPTIONS INDEX ALL
  EXPERIMENTAL SCRUB TABLE ... WITH OPTIONS INDEX (<index>...)
  EXPERIMENTAL SCRUB TABLE ... WITH OPTIONS CONSTRAINT ALL
  EXPERIMENTAL SCRUB TABLE ... WITH OPTIONS CONSTRAINT (<constraint>...)
  EXPERIMENTAL SCRUB TABLE ... WITH OPTIONS PHYSICAL
`,
		//line sql.y: 3039
		SeeAlso: `SCRUB DATABASE, SRUB
`,
	},
	//line sql.y: 3094
	`SET CLUSTER SETTING`: {
		ShortDescription: `change a cluster setting`,
		//line sql.y: 3095
		Category: hCfg,
		//line sql.y: 3096
		Text: `SET CLUSTER SETTING <var> { TO | = } <value>
`,
		//line sql.y: 3097
		SeeAlso: `SHOW CLUSTER SETTING, RESET CLUSTER SETTING, SET SESSION,
WEBDOCS/cluster-settings.html
`,
	},
	//line sql.y: 3118
	`SET SESSION`: {
		ShortDescription: `change a session variable`,
		//line sql.y: 3119
		Category: hCfg,
		//line sql.y: 3120
		Text: `
SET [SESSION] <var> { TO | = } <values...>
SET [SESSION] TIME ZONE <tz>
SET [SESSION] CHARACTERISTICS AS TRANSACTION ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }
SET [SESSION] TRACING { TO | = } { on | off | cluster | local | kv | results } [,...]

`,
		//line sql.y: 3126
		SeeAlso: `SHOW SESSION, RESET, DISCARD, SHOW, SET CLUSTER SETTING, SET TRANSACTION,
WEBDOCS/set-vars.html
`,
	},
	//line sql.y: 3143
	`SET TRANSACTION`: {
		ShortDescription: `configure the transaction settings`,
		//line sql.y: 3144
		Category: hTxn,
		//line sql.y: 3145
		Text: `
SET [SESSION] TRANSACTION <txnparameters...>

Transaction parameters:
   ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }
   PRIORITY { LOW | NORMAL | HIGH }

`,
		//line sql.y: 3152
		SeeAlso: `SHOW TRANSACTION, SET SESSION,
WEBDOCS/set-transaction.html
`,
	},
	//line sql.y: 3344
	`SHOW`: {
		//line sql.y: 3345
		Category: hGroup,
		//line sql.y: 3346
		Text: `
SHOW BACKUP, SHOW CLUSTER SETTING, SHOW COLUMNS, SHOW CONSTRAINTS,
SHOW CREATE, SHOW DATABASES, SHOW HISTOGRAM, SHOW INDEXES, SHOW
PARTITIONS, SHOW JOBS, SHOW QUERIES, SHOW RANGE, SHOW RANGES,
SHOW ROLES, SHOW SCHEMAS, SHOW SEQUENCES, SHOW SESSION, SHOW SESSIONS,
SHOW STATISTICS, SHOW SYNTAX, SHOW TABLES, SHOW TRACE SHOW TRANSACTION, SHOW USERS
`,
	},
	//line sql.y: 3382
	`SHOW SESSION`: {
		ShortDescription: `display session variables`,
		//line sql.y: 3383
		Category: hCfg,
		//line sql.y: 3384
		Text: `SHOW [SESSION] { <var> | ALL }
`,
		//line sql.y: 3385
		SeeAlso: `WEBDOCS/show-vars.html
`,
	},
	//line sql.y: 3406
	`SHOW STATISTICS`: {
		ShortDescription: `display table statistics (experimental)`,
		//line sql.y: 3407
		Category: hExperimental,
		//line sql.y: 3408
		Text: `SHOW STATISTICS [USING JSON] FOR TABLE <table_name>

Returns the available statistics for a table.
The statistics can include a histogram ID, which can
be used with SHOW HISTOGRAM.
If USING JSON is specified, the statistics and histograms
are encoded in JSON format.
`,
		//line sql.y: 3415
		SeeAlso: `SHOW HISTOGRAM
`,
	},
	//line sql.y: 3428
	`SHOW HISTOGRAM`: {
		ShortDescription: `display histogram (experimental)`,
		//line sql.y: 3429
		Category: hExperimental,
		//line sql.y: 3430
		Text: `SHOW HISTOGRAM <histogram_id>

Returns the data in the histogram with the
given ID (as returned by SHOW STATISTICS).
`,
		//line sql.y: 3434
		SeeAlso: `SHOW STATISTICS
`,
	},
	//line sql.y: 3447
	`SHOW BACKUP`: {
		ShortDescription: `list backup contents`,
		//line sql.y: 3448
		Category: hCCL,
		//line sql.y: 3449
		Text: `SHOW BACKUP [SCHEMAS|FILES|RANGES] <location>
`,
		//line sql.y: 3450
		SeeAlso: `WEBDOCS/show-backup.html
`,
	},
	//line sql.y: 3485
	`SHOW CLUSTER SETTING`: {
		ShortDescription: `display cluster settings`,
		//line sql.y: 3486
		Category: hCfg,
		//line sql.y: 3487
		Text: `
SHOW CLUSTER SETTING <var>
SHOW [ PUBLIC | ALL ] CLUSTER SETTINGS
`,
		//line sql.y: 3490
		SeeAlso: `WEBDOCS/cluster-settings.html
`,
	},
	//line sql.y: 3516
	`SHOW COLUMNS`: {
		ShortDescription: `list columns in relation`,
		//line sql.y: 3517
		Category: hDDL,
		//line sql.y: 3518
		Text: `SHOW COLUMNS FROM <tablename>
`,
		//line sql.y: 3519
		SeeAlso: `WEBDOCS/show-columns.html
`,
	},
	//line sql.y: 3527
	`SHOW PARTITIONS`: {
		ShortDescription: `list partition information`,
		//line sql.y: 3528
		Category: hDDL,
		//line sql.y: 3529
		Text: `SHOW PARTITIONS FROM { TABLE <table> | INDEX <index> | DATABASE <database> }
`,
		//line sql.y: 3530
		SeeAlso: `WEBDOCS/show-partitions.html
`,
	},
	//line sql.y: 3550
	`SHOW DATABASES`: {
		ShortDescription: `list databases`,
		//line sql.y: 3551
		Category: hDDL,
		//line sql.y: 3552
		Text: `SHOW DATABASES
`,
		//line sql.y: 3553
		SeeAlso: `WEBDOCS/show-databases.html
`,
	},
	//line sql.y: 3561
	`SHOW GRANTS`: {
		ShortDescription: `list grants`,
		//line sql.y: 3562
		Category: hPriv,
		//line sql.y: 3563
		Text: `
Show privilege grants:
  SHOW GRANTS [ON <targets...>] [FOR <users...>]
Show role grants:
  SHOW GRANTS ON ROLE [<roles...>] [FOR <grantees...>]

`,
		//line sql.y: 3569
		SeeAlso: `WEBDOCS/show-grants.html
`,
	},
	//line sql.y: 3582
	`SHOW INDEXES`: {
		ShortDescription: `list indexes`,
		//line sql.y: 3583
		Category: hDDL,
		//line sql.y: 3584
		Text: `SHOW INDEXES FROM { <tablename> | DATABASE <database_name> } [WITH COMMENT]
`,
		//line sql.y: 3585
		SeeAlso: `WEBDOCS/show-index.html
`,
	},
	//line sql.y: 3615
	`SHOW CONSTRAINTS`: {
		ShortDescription: `list constraints`,
		//line sql.y: 3616
		Category: hDDL,
		//line sql.y: 3617
		Text: `SHOW CONSTRAINTS FROM <tablename>
`,
		//line sql.y: 3618
		SeeAlso: `WEBDOCS/show-constraints.html
`,
	},
	//line sql.y: 3631
	`SHOW QUERIES`: {
		ShortDescription: `list running queries`,
		//line sql.y: 3632
		Category: hMisc,
		//line sql.y: 3633
		Text: `SHOW [ALL] [CLUSTER | LOCAL] QUERIES
`,
		//line sql.y: 3634
		SeeAlso: `CANCEL QUERIES
`,
	},
	//line sql.y: 3655
	`SHOW JOBS`: {
		ShortDescription: `list background jobs`,
		//line sql.y: 3656
		Category: hMisc,
		//line sql.y: 3657
		Text: `
SHOW [AUTOMATIC] JOBS
SHOW JOB <jobid>
`,
		//line sql.y: 3660
		SeeAlso: `CANCEL JOBS, PAUSE JOBS, RESUME JOBS
`,
	},
	//line sql.y: 3700
	`SHOW TRACE`: {
		ShortDescription: `display an execution trace`,
		//line sql.y: 3701
		Category: hMisc,
		//line sql.y: 3702
		Text: `
SHOW [COMPACT] [KV] TRACE FOR SESSION
`,
		//line sql.y: 3704
		SeeAlso: `EXPLAIN
`,
	},
	//line sql.y: 3727
	`SHOW SESSIONS`: {
		ShortDescription: `list open client sessions`,
		//line sql.y: 3728
		Category: hMisc,
		//line sql.y: 3729
		Text: `SHOW [ALL] [CLUSTER | LOCAL] SESSIONS
`,
		//line sql.y: 3730
		SeeAlso: `CANCEL SESSIONS
`,
	},
	//line sql.y: 3743
	`SHOW TABLES`: {
		ShortDescription: `list tables`,
		//line sql.y: 3744
		Category: hDDL,
		//line sql.y: 3745
		Text: `SHOW TABLES [FROM <databasename> [ . <schemaname> ] ] [WITH COMMENT]
`,
		//line sql.y: 3746
		SeeAlso: `WEBDOCS/show-tables.html
`,
	},
	//line sql.y: 3778
	`SHOW SCHEMAS`: {
		ShortDescription: `list schemas`,
		//line sql.y: 3779
		Category: hDDL,
		//line sql.y: 3780
		Text: `SHOW SCHEMAS [FROM <databasename> ]
`,
	},
	//line sql.y: 3792
	`SHOW SEQUENCES`: {
		ShortDescription: `list sequences`,
		//line sql.y: 3793
		Category: hDDL,
		//line sql.y: 3794
		Text: `SHOW SEQUENCES [FROM <databasename> ]
`,
	},
	//line sql.y: 3806
	`SHOW SYNTAX`: {
		ShortDescription: `analyze SQL syntax`,
		//line sql.y: 3807
		Category: hMisc,
		//line sql.y: 3808
		Text: `SHOW SYNTAX <string>
`,
	},
	//line sql.y: 3817
	`SHOW TRANSACTION`: {
		ShortDescription: `display current transaction properties`,
		//line sql.y: 3818
		Category: hCfg,
		//line sql.y: 3819
		Text: `SHOW TRANSACTION {ISOLATION LEVEL | PRIORITY | STATUS}
`,
		//line sql.y: 3820
		SeeAlso: `WEBDOCS/show-transaction.html
`,
	},
	//line sql.y: 3839
	`SHOW CREATE`: {
		ShortDescription: `display the CREATE statement for a table, sequence or view`,
		//line sql.y: 3840
		Category: hDDL,
		//line sql.y: 3841
		Text: `SHOW CREATE [ TABLE | SEQUENCE | VIEW ] <tablename>
`,
		//line sql.y: 3842
		SeeAlso: `WEBDOCS/show-create-table.html
`,
	},
	//line sql.y: 3860
	`SHOW USERS`: {
		ShortDescription: `list defined users`,
		//line sql.y: 3861
		Category: hPriv,
		//line sql.y: 3862
		Text: `SHOW USERS
`,
		//line sql.y: 3863
		SeeAlso: `CREATE USER, DROP USER, WEBDOCS/show-users.html
`,
	},
	//line sql.y: 3871
	`SHOW ROLES`: {
		ShortDescription: `list defined roles`,
		//line sql.y: 3872
		Category: hPriv,
		//line sql.y: 3873
		Text: `SHOW ROLES
`,
		//line sql.y: 3874
		SeeAlso: `CREATE ROLE, DROP ROLE
`,
	},
	//line sql.y: 3930
	`SHOW RANGE`: {
		ShortDescription: `show range information for a row`,
		//line sql.y: 3931
		Category: hMisc,
		//line sql.y: 3932
		Text: `
SHOW RANGE FROM TABLE <tablename> FOR ROW (row, value, ...)
SHOW RANGE FROM INDEX [ <tablename> @ ] <indexname> FOR ROW (row, value, ...)
`,
	},
	//line sql.y: 3953
	`SHOW RANGES`: {
		ShortDescription: `list ranges`,
		//line sql.y: 3954
		Category: hMisc,
		//line sql.y: 3955
		Text: `
SHOW RANGES FROM TABLE <tablename>
SHOW RANGES FROM INDEX [ <tablename> @ ] <indexname>
`,
	},
	//line sql.y: 4192
	`PAUSE JOBS`: {
		ShortDescription: `pause background jobs`,
		//line sql.y: 4193
		Category: hMisc,
		//line sql.y: 4194
		Text: `
PAUSE JOBS <selectclause>
PAUSE JOB <jobid>
`,
		//line sql.y: 4197
		SeeAlso: `SHOW JOBS, CANCEL JOBS, RESUME JOBS
`,
	},
	//line sql.y: 4214
	`CREATE TABLE`: {
		ShortDescription: `create a new table`,
		//line sql.y: 4215
		Category: hDDL,
		//line sql.y: 4216
		Text: `
CREATE [[GLOBAL | LOCAL] {TEMPORARY | TEMP}] TABLE [IF NOT EXISTS] <tablename> ( <elements...> ) [<interleave>]
CREATE [[GLOBAL | LOCAL] {TEMPORARY | TEMP}] TABLE [IF NOT EXISTS] <tablename> [( <colnames...> )] AS <source>

Table elements:
   <name> <type> [<qualifiers...>]
   [UNIQUE | INVERTED] INDEX [<name>] ( <colname> [ASC | DESC] [, ...] )
                           [STORING ( <colnames...> )] [<interleave>]
   FAMILY [<name>] ( <colnames...> )
   [CONSTRAINT <name>] <constraint>

Table constraints:
   PRIMARY KEY ( <colnames...> )
   FOREIGN KEY ( <colnames...> ) REFERENCES <tablename> [( <colnames...> )] [ON DELETE {NO ACTION | RESTRICT}] [ON UPDATE {NO ACTION | RESTRICT}]
   UNIQUE ( <colnames... ) [STORING ( <colnames...> )] [<interleave>]
   CHECK ( <expr> )

Column qualifiers:
  [CONSTRAINT <constraintname>] {NULL | NOT NULL | UNIQUE | PRIMARY KEY | CHECK (<expr>) | DEFAULT <expr>}
  FAMILY <familyname>, CREATE [IF NOT EXISTS] FAMILY [<familyname>]
  REFERENCES <tablename> [( <colnames...> )] [ON DELETE {NO ACTION | RESTRICT}] [ON UPDATE {NO ACTION | RESTRICT}]
  COLLATE <collationname>
  AS ( <expr> ) STORED

Interleave clause:
   INTERLEAVE IN PARENT <tablename> ( <colnames...> ) [CASCADE | RESTRICT]

`,
		//line sql.y: 4243
		SeeAlso: `SHOW TABLES, CREATE VIEW, SHOW CREATE,
WEBDOCS/create-table.html
WEBDOCS/create-table-as.html
`,
	},
	//line sql.y: 4981
	`CREATE SEQUENCE`: {
		ShortDescription: `create a new sequence`,
		//line sql.y: 4982
		Category: hDDL,
		//line sql.y: 4983
		Text: `
CREATE [TEMPORARY | TEMP] SEQUENCE <seqname>
  [INCREMENT <increment>]
  [MINVALUE <minvalue> | NO MINVALUE]
  [MAXVALUE <maxvalue> | NO MAXVALUE]
  [START [WITH] <start>]
  [CACHE <cache>]
  [NO CYCLE]
  [VIRTUAL]

`,
		//line sql.y: 4993
		SeeAlso: `CREATE TABLE
`,
	},
	//line sql.y: 5058
	`TRUNCATE`: {
		ShortDescription: `empty one or more tables`,
		//line sql.y: 5059
		Category: hDML,
		//line sql.y: 5060
		Text: `TRUNCATE [TABLE] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 5061
		SeeAlso: `WEBDOCS/truncate.html
`,
	},
	//line sql.y: 5069
	`CREATE USER`: {
		ShortDescription: `define a new user`,
		//line sql.y: 5070
		Category: hPriv,
		//line sql.y: 5071
		Text: `CREATE USER [IF NOT EXISTS] <name> [ [WITH] PASSWORD <passwd> ]
`,
		//line sql.y: 5072
		SeeAlso: `DROP USER, SHOW USERS, WEBDOCS/create-user.html
`,
	},
	//line sql.y: 5101
	`CREATE ROLE`: {
		ShortDescription: `define a new role`,
		//line sql.y: 5102
		Category: hPriv,
		//line sql.y: 5103
		Text: `CREATE ROLE [IF NOT EXISTS] <name>
`,
		//line sql.y: 5104
		SeeAlso: `DROP ROLE, SHOW ROLES
`,
	},
	//line sql.y: 5122
	`CREATE VIEW`: {
		ShortDescription: `create a new view`,
		//line sql.y: 5123
		Category: hDDL,
		//line sql.y: 5124
		Text: `CREATE [TEMPORARY | TEMP] VIEW <viewname> [( <colnames...> )] AS <source>
`,
		//line sql.y: 5125
		SeeAlso: `CREATE TABLE, SHOW CREATE, WEBDOCS/create-view.html
`,
	},
	//line sql.y: 5160
	`CREATE INDEX`: {
		ShortDescription: `create a new index`,
		//line sql.y: 5161
		Category: hDDL,
		//line sql.y: 5162
		Text: `
CREATE [UNIQUE | INVERTED] INDEX [IF NOT EXISTS] [<idxname>]
       ON <tablename> ( <colname> [ASC | DESC] [, ...] )
       [STORING ( <colnames...> )] [<interleave>]

Interleave clause:
   INTERLEAVE IN PARENT <tablename> ( <colnames...> ) [CASCADE | RESTRICT]

`,
		//line sql.y: 5170
		SeeAlso: `CREATE TABLE, SHOW INDEXES, SHOW CREATE,
WEBDOCS/create-index.html
`,
	},
	//line sql.y: 5399
	`RELEASE`: {
		ShortDescription: `complete a retryable block`,
		//line sql.y: 5400
		Category: hTxn,
		//line sql.y: 5401
		Text: `RELEASE [SAVEPOINT] cockroach_restart
`,
		//line sql.y: 5402
		SeeAlso: `SAVEPOINT, WEBDOCS/savepoint.html
`,
	},
	//line sql.y: 5410
	`RESUME JOBS`: {
		ShortDescription: `resume background jobs`,
		//line sql.y: 5411
		Category: hMisc,
		//line sql.y: 5412
		Text: `
RESUME JOBS <selectclause>
RESUME JOB <jobid>
`,
		//line sql.y: 5415
		SeeAlso: `SHOW JOBS, CANCEL JOBS, PAUSE JOBS
`,
	},
	//line sql.y: 5432
	`SAVEPOINT`: {
		ShortDescription: `start a retryable block`,
		//line sql.y: 5433
		Category: hTxn,
		//line sql.y: 5434
		Text: `SAVEPOINT cockroach_restart
`,
		//line sql.y: 5435
		SeeAlso: `RELEASE, WEBDOCS/savepoint.html
`,
	},
	//line sql.y: 5450
	`BEGIN`: {
		ShortDescription: `start a transaction`,
		//line sql.y: 5451
		Category: hTxn,
		//line sql.y: 5452
		Text: `
BEGIN [TRANSACTION] [ <txnparameter> [[,] ...] ]
START TRANSACTION [ <txnparameter> [[,] ...] ]

Transaction parameters:
   ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }
   PRIORITY { LOW | NORMAL | HIGH }

`,
		//line sql.y: 5460
		SeeAlso: `COMMIT, ROLLBACK, WEBDOCS/begin-transaction.html
`,
	},
	//line sql.y: 5473
	`COMMIT`: {
		ShortDescription: `commit the current transaction`,
		//line sql.y: 5474
		Category: hTxn,
		//line sql.y: 5475
		Text: `
COMMIT [TRANSACTION]
END [TRANSACTION]
`,
		//line sql.y: 5478
		SeeAlso: `BEGIN, ROLLBACK, WEBDOCS/commit-transaction.html
`,
	},
	//line sql.y: 5502
	`ROLLBACK`: {
		ShortDescription: `abort the current transaction`,
		//line sql.y: 5503
		Category: hTxn,
		//line sql.y: 5504
		Text: `ROLLBACK [TRANSACTION] [TO [SAVEPOINT] cockroach_restart]
`,
		//line sql.y: 5505
		SeeAlso: `BEGIN, COMMIT, SAVEPOINT, WEBDOCS/rollback-transaction.html
`,
	},
	//line sql.y: 5623
	`CREATE DATABASE`: {
		ShortDescription: `create a new database`,
		//line sql.y: 5624
		Category: hDDL,
		//line sql.y: 5625
		Text: `CREATE DATABASE [IF NOT EXISTS] <name>
`,
		//line sql.y: 5626
		SeeAlso: `WEBDOCS/create-database.html
`,
	},
	//line sql.y: 5695
	`INSERT`: {
		ShortDescription: `create new rows in a table`,
		//line sql.y: 5696
		Category: hDML,
		//line sql.y: 5697
		Text: `
INSERT INTO <tablename> [[AS] <name>] [( <colnames...> )]
       <selectclause>
       [ON CONFLICT [( <colnames...> )] {DO UPDATE SET ... [WHERE <expr>] | DO NOTHING}]
       [RETURNING <exprs...>]
`,
		//line sql.y: 5702
		SeeAlso: `UPSERT, UPDATE, DELETE, WEBDOCS/insert.html
`,
	},
	//line sql.y: 5721
	`UPSERT`: {
		ShortDescription: `create or replace rows in a table`,
		//line sql.y: 5722
		Category: hDML,
		//line sql.y: 5723
		Text: `
UPSERT INTO <tablename> [AS <name>] [( <colnames...> )]
       <selectclause>
       [RETURNING <exprs...>]
`,
		//line sql.y: 5727
		SeeAlso: `INSERT, UPDATE, DELETE, WEBDOCS/upsert.html
`,
	},
	//line sql.y: 5838
	`UPDATE`: {
		ShortDescription: `update rows of a table`,
		//line sql.y: 5839
		Category: hDML,
		//line sql.y: 5840
		Text: `
UPDATE <tablename> [[AS] <name>]
       SET ...
       [WHERE <expr>]
       [ORDER BY <exprs...>]
       [LIMIT <expr>]
       [RETURNING <exprs...>]
`,
		//line sql.y: 5847
		SeeAlso: `INSERT, UPSERT, DELETE, WEBDOCS/update.html
`,
	},
	//line sql.y: 6072
	`<SELECTCLAUSE>`: {
		ShortDescription: `access tabular data`,
		//line sql.y: 6073
		Category: hDML,
		//line sql.y: 6074
		Text: `
Select clause:
  TABLE <tablename>
  VALUES ( <exprs...> ) [ , ... ]
  SELECT ... [ { INTERSECT | UNION | EXCEPT } [ ALL | DISTINCT ] <selectclause> ]
`,
	},
	//line sql.y: 6085
	`SELECT`: {
		ShortDescription: `retrieve rows from a data source and compute a result`,
		//line sql.y: 6086
		Category: hDML,
		//line sql.y: 6087
		Text: `
SELECT [DISTINCT [ ON ( <expr> [ , ... ] ) ] ]
       { <expr> [[AS] <name>] | [ [<dbname>.] <tablename>. ] * } [, ...]
       [ FROM <source> ]
       [ WHERE <expr> ]
       [ GROUP BY <expr> [ , ... ] ]
       [ HAVING <expr> ]
       [ WINDOW <name> AS ( <definition> ) ]
       [ { UNION | INTERSECT | EXCEPT } [ ALL | DISTINCT ] <selectclause> ]
       [ ORDER BY <expr> [ ASC | DESC ] [, ...] ]
       [ LIMIT { <expr> | ALL } ]
       [ OFFSET <expr> [ ROW | ROWS ] ]
`,
		//line sql.y: 6099
		SeeAlso: `WEBDOCS/select-clause.html
`,
	},
	//line sql.y: 6174
	`TABLE`: {
		ShortDescription: `select an entire table`,
		//line sql.y: 6175
		Category: hDML,
		//line sql.y: 6176
		Text: `TABLE <tablename>
`,
		//line sql.y: 6177
		SeeAlso: `SELECT, VALUES, WEBDOCS/table-expressions.html
`,
	},
	//line sql.y: 6489
	`VALUES`: {
		ShortDescription: `select a given set of values`,
		//line sql.y: 6490
		Category: hDML,
		//line sql.y: 6491
		Text: `VALUES ( <exprs...> ) [, ...]
`,
		//line sql.y: 6492
		SeeAlso: `SELECT, TABLE, WEBDOCS/table-expressions.html
`,
	},
	//line sql.y: 6601
	`<SOURCE>`: {
		ShortDescription: `define a data source for SELECT`,
		//line sql.y: 6602
		Category: hDML,
		//line sql.y: 6603
		Text: `
Data sources:
  <tablename> [ @ { <idxname> | <indexflags> } ]
  <tablefunc> ( <exprs...> )
  ( { <selectclause> | <source> } )
  <source> [AS] <alias> [( <colnames...> )]
  <source> [ <jointype> ] JOIN <source> ON <expr>
  <source> [ <jointype> ] JOIN <source> USING ( <colnames...> )
  <source> NATURAL [ <jointype> ] JOIN <source>
  <source> CROSS JOIN <source>
  <source> WITH ORDINALITY
  '[' EXPLAIN ... ']'
  '[' SHOW ... ']'

Index flags:
  '{' FORCE_INDEX = <idxname> [, ...] '}'
  '{' NO_INDEX_JOIN [, ...] '}'
  '{' IGNORE_FOREIGN_KEYS [, ...] '}'

Join types:
  { INNER | { LEFT | RIGHT | FULL } [OUTER] } [ { HASH | MERGE | LOOKUP } ]

`,
		//line sql.y: 6625
		SeeAlso: `WEBDOCS/table-expressions.html
`,
	},
}
