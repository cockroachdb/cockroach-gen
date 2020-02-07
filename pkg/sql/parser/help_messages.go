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
	//line sql.y: 1913
	`IMPORT`: {
		ShortDescription: `load data from file in a distributed manner`,
		//line sql.y: 1914
		Category: hCCL,
		//line sql.y: 1915
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
		//line sql.y: 1943
		SeeAlso: `CREATE TABLE
`,
	},
	//line sql.y: 1987
	`EXPORT`: {
		ShortDescription: `export data to file in a distributed manner`,
		//line sql.y: 1988
		Category: hCCL,
		//line sql.y: 1989
		Text: `
EXPORT INTO <format> <datafile> [WITH <option> [= value] [,...]] FROM <query>

Formats:
   CSV

Options:
   delimiter = '...'   [CSV-specific]

`,
		//line sql.y: 1998
		SeeAlso: `SELECT
`,
	},
	//line sql.y: 2092
	`CANCEL`: {
		//line sql.y: 2093
		Category: hGroup,
		//line sql.y: 2094
		Text: `CANCEL JOBS, CANCEL QUERIES, CANCEL SESSIONS
`,
	},
	//line sql.y: 2101
	`CANCEL JOBS`: {
		ShortDescription: `cancel background jobs`,
		//line sql.y: 2102
		Category: hMisc,
		//line sql.y: 2103
		Text: `
CANCEL JOBS <selectclause>
CANCEL JOB <jobid>
`,
		//line sql.y: 2106
		SeeAlso: `SHOW JOBS, PAUSE JOBS, RESUME JOBS
`,
	},
	//line sql.y: 2124
	`CANCEL QUERIES`: {
		ShortDescription: `cancel running queries`,
		//line sql.y: 2125
		Category: hMisc,
		//line sql.y: 2126
		Text: `
CANCEL QUERIES [IF EXISTS] <selectclause>
CANCEL QUERY [IF EXISTS] <expr>
`,
		//line sql.y: 2129
		SeeAlso: `SHOW QUERIES
`,
	},
	//line sql.y: 2160
	`CANCEL SESSIONS`: {
		ShortDescription: `cancel open sessions`,
		//line sql.y: 2161
		Category: hMisc,
		//line sql.y: 2162
		Text: `
CANCEL SESSIONS [IF EXISTS] <selectclause>
CANCEL SESSION [IF EXISTS] <sessionid>
`,
		//line sql.y: 2165
		SeeAlso: `SHOW SESSIONS
`,
	},
	//line sql.y: 2235
	`CREATE`: {
		//line sql.y: 2236
		Category: hGroup,
		//line sql.y: 2237
		Text: `
CREATE DATABASE, CREATE TABLE, CREATE INDEX, CREATE TABLE AS,
CREATE USER, CREATE VIEW, CREATE SEQUENCE, CREATE STATISTICS,
CREATE ROLE
`,
	},
	//line sql.y: 2318
	`CREATE STATISTICS`: {
		ShortDescription: `create a new table statistic`,
		//line sql.y: 2319
		Category: hMisc,
		//line sql.y: 2320
		Text: `
CREATE STATISTICS <statisticname>
  [ON <colname> [, ...]]
  FROM <tablename> [AS OF SYSTEM TIME <expr>]
`,
	},
	//line sql.y: 2463
	`DELETE`: {
		ShortDescription: `delete rows from a table`,
		//line sql.y: 2464
		Category: hDML,
		//line sql.y: 2465
		Text: `DELETE FROM <tablename> [WHERE <expr>]
              [ORDER BY <exprs...>]
              [LIMIT <expr>]
              [RETURNING <exprs...>]
`,
		//line sql.y: 2469
		SeeAlso: `WEBDOCS/delete.html
`,
	},
	//line sql.y: 2489
	`DISCARD`: {
		ShortDescription: `reset the session to its initial state`,
		//line sql.y: 2490
		Category: hCfg,
		//line sql.y: 2491
		Text: `DISCARD ALL
`,
	},
	//line sql.y: 2503
	`DROP`: {
		//line sql.y: 2504
		Category: hGroup,
		//line sql.y: 2505
		Text: `
DROP DATABASE, DROP INDEX, DROP TABLE, DROP VIEW, DROP SEQUENCE,
DROP USER, DROP ROLE
`,
	},
	//line sql.y: 2522
	`DROP VIEW`: {
		ShortDescription: `remove a view`,
		//line sql.y: 2523
		Category: hDDL,
		//line sql.y: 2524
		Text: `DROP VIEW [IF EXISTS] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 2525
		SeeAlso: `WEBDOCS/drop-index.html
`,
	},
	//line sql.y: 2537
	`DROP SEQUENCE`: {
		ShortDescription: `remove a sequence`,
		//line sql.y: 2538
		Category: hDDL,
		//line sql.y: 2539
		Text: `DROP SEQUENCE [IF EXISTS] <sequenceName> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 2540
		SeeAlso: `DROP
`,
	},
	//line sql.y: 2552
	`DROP TABLE`: {
		ShortDescription: `remove a table`,
		//line sql.y: 2553
		Category: hDDL,
		//line sql.y: 2554
		Text: `DROP TABLE [IF EXISTS] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 2555
		SeeAlso: `WEBDOCS/drop-table.html
`,
	},
	//line sql.y: 2567
	`DROP INDEX`: {
		ShortDescription: `remove an index`,
		//line sql.y: 2568
		Category: hDDL,
		//line sql.y: 2569
		Text: `DROP INDEX [IF EXISTS] <idxname> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 2570
		SeeAlso: `WEBDOCS/drop-index.html
`,
	},
	//line sql.y: 2590
	`DROP DATABASE`: {
		ShortDescription: `remove a database`,
		//line sql.y: 2591
		Category: hDDL,
		//line sql.y: 2592
		Text: `DROP DATABASE [IF EXISTS] <databasename> [CASCADE | RESTRICT]
`,
		//line sql.y: 2593
		SeeAlso: `WEBDOCS/drop-database.html
`,
	},
	//line sql.y: 2613
	`DROP USER`: {
		ShortDescription: `remove a user`,
		//line sql.y: 2614
		Category: hPriv,
		//line sql.y: 2615
		Text: `DROP USER [IF EXISTS] <user> [, ...]
`,
		//line sql.y: 2616
		SeeAlso: `CREATE USER, SHOW USERS
`,
	},
	//line sql.y: 2628
	`DROP ROLE`: {
		ShortDescription: `remove a role`,
		//line sql.y: 2629
		Category: hPriv,
		//line sql.y: 2630
		Text: `DROP ROLE [IF EXISTS] <role> [, ...]
`,
		//line sql.y: 2631
		SeeAlso: `CREATE ROLE, SHOW ROLES
`,
	},
	//line sql.y: 2655
	`EXPLAIN`: {
		ShortDescription: `show the logical plan of a query`,
		//line sql.y: 2656
		Category: hMisc,
		//line sql.y: 2657
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
		//line sql.y: 2670
		SeeAlso: `WEBDOCS/explain.html
`,
	},
	//line sql.y: 2753
	`PREPARE`: {
		ShortDescription: `prepare a statement for later execution`,
		//line sql.y: 2754
		Category: hMisc,
		//line sql.y: 2755
		Text: `PREPARE <name> [ ( <types...> ) ] AS <query>
`,
		//line sql.y: 2756
		SeeAlso: `EXECUTE, DEALLOCATE, DISCARD
`,
	},
	//line sql.y: 2787
	`EXECUTE`: {
		ShortDescription: `execute a statement prepared previously`,
		//line sql.y: 2788
		Category: hMisc,
		//line sql.y: 2789
		Text: `EXECUTE <name> [ ( <exprs...> ) ]
`,
		//line sql.y: 2790
		SeeAlso: `PREPARE, DEALLOCATE, DISCARD
`,
	},
	//line sql.y: 2820
	`DEALLOCATE`: {
		ShortDescription: `remove a prepared statement`,
		//line sql.y: 2821
		Category: hMisc,
		//line sql.y: 2822
		Text: `DEALLOCATE [PREPARE] { <name> | ALL }
`,
		//line sql.y: 2823
		SeeAlso: `PREPARE, EXECUTE, DISCARD
`,
	},
	//line sql.y: 2843
	`GRANT`: {
		ShortDescription: `define access privileges and role memberships`,
		//line sql.y: 2844
		Category: hPriv,
		//line sql.y: 2845
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
		//line sql.y: 2858
		SeeAlso: `REVOKE, WEBDOCS/grant.html
`,
	},
	//line sql.y: 2874
	`REVOKE`: {
		ShortDescription: `remove access privileges and role memberships`,
		//line sql.y: 2875
		Category: hPriv,
		//line sql.y: 2876
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
		//line sql.y: 2889
		SeeAlso: `GRANT, WEBDOCS/revoke.html
`,
	},
	//line sql.y: 2943
	`RESET`: {
		ShortDescription: `reset a session variable to its default value`,
		//line sql.y: 2944
		Category: hCfg,
		//line sql.y: 2945
		Text: `RESET [SESSION] <var>
`,
		//line sql.y: 2946
		SeeAlso: `RESET CLUSTER SETTING, WEBDOCS/set-vars.html
`,
	},
	//line sql.y: 2958
	`RESET CLUSTER SETTING`: {
		ShortDescription: `reset a cluster setting to its default value`,
		//line sql.y: 2959
		Category: hCfg,
		//line sql.y: 2960
		Text: `RESET CLUSTER SETTING <var>
`,
		//line sql.y: 2961
		SeeAlso: `SET CLUSTER SETTING, RESET
`,
	},
	//line sql.y: 2970
	`USE`: {
		ShortDescription: `set the current database`,
		//line sql.y: 2971
		Category: hCfg,
		//line sql.y: 2972
		Text: `USE <dbname>

"USE <dbname>" is an alias for "SET [SESSION] database = <dbname>".
`,
		//line sql.y: 2975
		SeeAlso: `SET SESSION, WEBDOCS/set-vars.html
`,
	},
	//line sql.y: 2996
	`SCRUB`: {
		ShortDescription: `run checks against databases or tables`,
		//line sql.y: 2997
		Category: hExperimental,
		//line sql.y: 2998
		Text: `
EXPERIMENTAL SCRUB TABLE <table> ...
EXPERIMENTAL SCRUB DATABASE <database>

The various checks that ca be run with SCRUB includes:
  - Physical table data (encoding)
  - Secondary index integrity
  - Constraint integrity (NOT NULL, CHECK, FOREIGN KEY, UNIQUE)
`,
		//line sql.y: 3006
		SeeAlso: `SCRUB TABLE, SCRUB DATABASE
`,
	},
	//line sql.y: 3012
	`SCRUB DATABASE`: {
		ShortDescription: `run scrub checks on a database`,
		//line sql.y: 3013
		Category: hExperimental,
		//line sql.y: 3014
		Text: `
EXPERIMENTAL SCRUB DATABASE <database>
                            [AS OF SYSTEM TIME <expr>]

All scrub checks will be run on the database. This includes:
  - Physical table data (encoding)
  - Secondary index integrity
  - Constraint integrity (NOT NULL, CHECK, FOREIGN KEY, UNIQUE)
`,
		//line sql.y: 3022
		SeeAlso: `SCRUB TABLE, SCRUB
`,
	},
	//line sql.y: 3030
	`SCRUB TABLE`: {
		ShortDescription: `run scrub checks on a table`,
		//line sql.y: 3031
		Category: hExperimental,
		//line sql.y: 3032
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
		//line sql.y: 3043
		SeeAlso: `SCRUB DATABASE, SRUB
`,
	},
	//line sql.y: 3098
	`SET CLUSTER SETTING`: {
		ShortDescription: `change a cluster setting`,
		//line sql.y: 3099
		Category: hCfg,
		//line sql.y: 3100
		Text: `SET CLUSTER SETTING <var> { TO | = } <value>
`,
		//line sql.y: 3101
		SeeAlso: `SHOW CLUSTER SETTING, RESET CLUSTER SETTING, SET SESSION,
WEBDOCS/cluster-settings.html
`,
	},
	//line sql.y: 3122
	`SET SESSION`: {
		ShortDescription: `change a session variable`,
		//line sql.y: 3123
		Category: hCfg,
		//line sql.y: 3124
		Text: `
SET [SESSION] <var> { TO | = } <values...>
SET [SESSION] TIME ZONE <tz>
SET [SESSION] CHARACTERISTICS AS TRANSACTION ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }
SET [SESSION] TRACING { TO | = } { on | off | cluster | local | kv | results } [,...]

`,
		//line sql.y: 3130
		SeeAlso: `SHOW SESSION, RESET, DISCARD, SHOW, SET CLUSTER SETTING, SET TRANSACTION,
WEBDOCS/set-vars.html
`,
	},
	//line sql.y: 3147
	`SET TRANSACTION`: {
		ShortDescription: `configure the transaction settings`,
		//line sql.y: 3148
		Category: hTxn,
		//line sql.y: 3149
		Text: `
SET [SESSION] TRANSACTION <txnparameters...>

Transaction parameters:
   ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }
   PRIORITY { LOW | NORMAL | HIGH }

`,
		//line sql.y: 3156
		SeeAlso: `SHOW TRANSACTION, SET SESSION,
WEBDOCS/set-transaction.html
`,
	},
	//line sql.y: 3348
	`SHOW`: {
		//line sql.y: 3349
		Category: hGroup,
		//line sql.y: 3350
		Text: `
SHOW BACKUP, SHOW CLUSTER SETTING, SHOW COLUMNS, SHOW CONSTRAINTS,
SHOW CREATE, SHOW DATABASES, SHOW HISTOGRAM, SHOW INDEXES, SHOW
PARTITIONS, SHOW JOBS, SHOW QUERIES, SHOW RANGE, SHOW RANGES,
SHOW ROLES, SHOW SCHEMAS, SHOW SEQUENCES, SHOW SESSION, SHOW SESSIONS,
SHOW STATISTICS, SHOW SYNTAX, SHOW TABLES, SHOW TRACE SHOW TRANSACTION, SHOW USERS
`,
	},
	//line sql.y: 3386
	`SHOW SESSION`: {
		ShortDescription: `display session variables`,
		//line sql.y: 3387
		Category: hCfg,
		//line sql.y: 3388
		Text: `SHOW [SESSION] { <var> | ALL }
`,
		//line sql.y: 3389
		SeeAlso: `WEBDOCS/show-vars.html
`,
	},
	//line sql.y: 3410
	`SHOW STATISTICS`: {
		ShortDescription: `display table statistics (experimental)`,
		//line sql.y: 3411
		Category: hExperimental,
		//line sql.y: 3412
		Text: `SHOW STATISTICS [USING JSON] FOR TABLE <table_name>

Returns the available statistics for a table.
The statistics can include a histogram ID, which can
be used with SHOW HISTOGRAM.
If USING JSON is specified, the statistics and histograms
are encoded in JSON format.
`,
		//line sql.y: 3419
		SeeAlso: `SHOW HISTOGRAM
`,
	},
	//line sql.y: 3432
	`SHOW HISTOGRAM`: {
		ShortDescription: `display histogram (experimental)`,
		//line sql.y: 3433
		Category: hExperimental,
		//line sql.y: 3434
		Text: `SHOW HISTOGRAM <histogram_id>

Returns the data in the histogram with the
given ID (as returned by SHOW STATISTICS).
`,
		//line sql.y: 3438
		SeeAlso: `SHOW STATISTICS
`,
	},
	//line sql.y: 3451
	`SHOW BACKUP`: {
		ShortDescription: `list backup contents`,
		//line sql.y: 3452
		Category: hCCL,
		//line sql.y: 3453
		Text: `SHOW BACKUP [SCHEMAS|FILES|RANGES] <location>
`,
		//line sql.y: 3454
		SeeAlso: `WEBDOCS/show-backup.html
`,
	},
	//line sql.y: 3493
	`SHOW CLUSTER SETTING`: {
		ShortDescription: `display cluster settings`,
		//line sql.y: 3494
		Category: hCfg,
		//line sql.y: 3495
		Text: `
SHOW CLUSTER SETTING <var>
SHOW [ PUBLIC | ALL ] CLUSTER SETTINGS
`,
		//line sql.y: 3498
		SeeAlso: `WEBDOCS/cluster-settings.html
`,
	},
	//line sql.y: 3524
	`SHOW COLUMNS`: {
		ShortDescription: `list columns in relation`,
		//line sql.y: 3525
		Category: hDDL,
		//line sql.y: 3526
		Text: `SHOW COLUMNS FROM <tablename>
`,
		//line sql.y: 3527
		SeeAlso: `WEBDOCS/show-columns.html
`,
	},
	//line sql.y: 3535
	`SHOW PARTITIONS`: {
		ShortDescription: `list partition information`,
		//line sql.y: 3536
		Category: hDDL,
		//line sql.y: 3537
		Text: `SHOW PARTITIONS FROM { TABLE <table> | INDEX <index> | DATABASE <database> }
`,
		//line sql.y: 3538
		SeeAlso: `WEBDOCS/show-partitions.html
`,
	},
	//line sql.y: 3558
	`SHOW DATABASES`: {
		ShortDescription: `list databases`,
		//line sql.y: 3559
		Category: hDDL,
		//line sql.y: 3560
		Text: `SHOW DATABASES
`,
		//line sql.y: 3561
		SeeAlso: `WEBDOCS/show-databases.html
`,
	},
	//line sql.y: 3569
	`SHOW GRANTS`: {
		ShortDescription: `list grants`,
		//line sql.y: 3570
		Category: hPriv,
		//line sql.y: 3571
		Text: `
Show privilege grants:
  SHOW GRANTS [ON <targets...>] [FOR <users...>]
Show role grants:
  SHOW GRANTS ON ROLE [<roles...>] [FOR <grantees...>]

`,
		//line sql.y: 3577
		SeeAlso: `WEBDOCS/show-grants.html
`,
	},
	//line sql.y: 3590
	`SHOW INDEXES`: {
		ShortDescription: `list indexes`,
		//line sql.y: 3591
		Category: hDDL,
		//line sql.y: 3592
		Text: `SHOW INDEXES FROM { <tablename> | DATABASE <database_name> } [WITH COMMENT]
`,
		//line sql.y: 3593
		SeeAlso: `WEBDOCS/show-index.html
`,
	},
	//line sql.y: 3623
	`SHOW CONSTRAINTS`: {
		ShortDescription: `list constraints`,
		//line sql.y: 3624
		Category: hDDL,
		//line sql.y: 3625
		Text: `SHOW CONSTRAINTS FROM <tablename>
`,
		//line sql.y: 3626
		SeeAlso: `WEBDOCS/show-constraints.html
`,
	},
	//line sql.y: 3639
	`SHOW QUERIES`: {
		ShortDescription: `list running queries`,
		//line sql.y: 3640
		Category: hMisc,
		//line sql.y: 3641
		Text: `SHOW [ALL] [CLUSTER | LOCAL] QUERIES
`,
		//line sql.y: 3642
		SeeAlso: `CANCEL QUERIES
`,
	},
	//line sql.y: 3663
	`SHOW JOBS`: {
		ShortDescription: `list background jobs`,
		//line sql.y: 3664
		Category: hMisc,
		//line sql.y: 3665
		Text: `
SHOW [AUTOMATIC] JOBS
SHOW JOB <jobid>
`,
		//line sql.y: 3668
		SeeAlso: `CANCEL JOBS, PAUSE JOBS, RESUME JOBS
`,
	},
	//line sql.y: 3708
	`SHOW TRACE`: {
		ShortDescription: `display an execution trace`,
		//line sql.y: 3709
		Category: hMisc,
		//line sql.y: 3710
		Text: `
SHOW [COMPACT] [KV] TRACE FOR SESSION
`,
		//line sql.y: 3712
		SeeAlso: `EXPLAIN
`,
	},
	//line sql.y: 3735
	`SHOW SESSIONS`: {
		ShortDescription: `list open client sessions`,
		//line sql.y: 3736
		Category: hMisc,
		//line sql.y: 3737
		Text: `SHOW [ALL] [CLUSTER | LOCAL] SESSIONS
`,
		//line sql.y: 3738
		SeeAlso: `CANCEL SESSIONS
`,
	},
	//line sql.y: 3751
	`SHOW TABLES`: {
		ShortDescription: `list tables`,
		//line sql.y: 3752
		Category: hDDL,
		//line sql.y: 3753
		Text: `SHOW TABLES [FROM <databasename> [ . <schemaname> ] ] [WITH COMMENT]
`,
		//line sql.y: 3754
		SeeAlso: `WEBDOCS/show-tables.html
`,
	},
	//line sql.y: 3786
	`SHOW SCHEMAS`: {
		ShortDescription: `list schemas`,
		//line sql.y: 3787
		Category: hDDL,
		//line sql.y: 3788
		Text: `SHOW SCHEMAS [FROM <databasename> ]
`,
	},
	//line sql.y: 3800
	`SHOW SEQUENCES`: {
		ShortDescription: `list sequences`,
		//line sql.y: 3801
		Category: hDDL,
		//line sql.y: 3802
		Text: `SHOW SEQUENCES [FROM <databasename> ]
`,
	},
	//line sql.y: 3814
	`SHOW SYNTAX`: {
		ShortDescription: `analyze SQL syntax`,
		//line sql.y: 3815
		Category: hMisc,
		//line sql.y: 3816
		Text: `SHOW SYNTAX <string>
`,
	},
	//line sql.y: 3825
	`SHOW TRANSACTION`: {
		ShortDescription: `display current transaction properties`,
		//line sql.y: 3826
		Category: hCfg,
		//line sql.y: 3827
		Text: `SHOW TRANSACTION {ISOLATION LEVEL | PRIORITY | STATUS}
`,
		//line sql.y: 3828
		SeeAlso: `WEBDOCS/show-transaction.html
`,
	},
	//line sql.y: 3847
	`SHOW CREATE`: {
		ShortDescription: `display the CREATE statement for a table, sequence or view`,
		//line sql.y: 3848
		Category: hDDL,
		//line sql.y: 3849
		Text: `SHOW CREATE [ TABLE | SEQUENCE | VIEW ] <tablename>
`,
		//line sql.y: 3850
		SeeAlso: `WEBDOCS/show-create-table.html
`,
	},
	//line sql.y: 3868
	`SHOW USERS`: {
		ShortDescription: `list defined users`,
		//line sql.y: 3869
		Category: hPriv,
		//line sql.y: 3870
		Text: `SHOW USERS
`,
		//line sql.y: 3871
		SeeAlso: `CREATE USER, DROP USER, WEBDOCS/show-users.html
`,
	},
	//line sql.y: 3879
	`SHOW ROLES`: {
		ShortDescription: `list defined roles`,
		//line sql.y: 3880
		Category: hPriv,
		//line sql.y: 3881
		Text: `SHOW ROLES
`,
		//line sql.y: 3882
		SeeAlso: `CREATE ROLE, DROP ROLE
`,
	},
	//line sql.y: 3938
	`SHOW RANGE`: {
		ShortDescription: `show range information for a row`,
		//line sql.y: 3939
		Category: hMisc,
		//line sql.y: 3940
		Text: `
SHOW RANGE FROM TABLE <tablename> FOR ROW (row, value, ...)
SHOW RANGE FROM INDEX [ <tablename> @ ] <indexname> FOR ROW (row, value, ...)
`,
	},
	//line sql.y: 3961
	`SHOW RANGES`: {
		ShortDescription: `list ranges`,
		//line sql.y: 3962
		Category: hMisc,
		//line sql.y: 3963
		Text: `
SHOW RANGES FROM TABLE <tablename>
SHOW RANGES FROM INDEX [ <tablename> @ ] <indexname>
`,
	},
	//line sql.y: 4200
	`PAUSE JOBS`: {
		ShortDescription: `pause background jobs`,
		//line sql.y: 4201
		Category: hMisc,
		//line sql.y: 4202
		Text: `
PAUSE JOBS <selectclause>
PAUSE JOB <jobid>
`,
		//line sql.y: 4205
		SeeAlso: `SHOW JOBS, CANCEL JOBS, RESUME JOBS
`,
	},
	//line sql.y: 4222
	`CREATE TABLE`: {
		ShortDescription: `create a new table`,
		//line sql.y: 4223
		Category: hDDL,
		//line sql.y: 4224
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
		//line sql.y: 4251
		SeeAlso: `SHOW TABLES, CREATE VIEW, SHOW CREATE,
WEBDOCS/create-table.html
WEBDOCS/create-table-as.html
`,
	},
	//line sql.y: 4989
	`CREATE SEQUENCE`: {
		ShortDescription: `create a new sequence`,
		//line sql.y: 4990
		Category: hDDL,
		//line sql.y: 4991
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
		//line sql.y: 5001
		SeeAlso: `CREATE TABLE
`,
	},
	//line sql.y: 5066
	`TRUNCATE`: {
		ShortDescription: `empty one or more tables`,
		//line sql.y: 5067
		Category: hDML,
		//line sql.y: 5068
		Text: `TRUNCATE [TABLE] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 5069
		SeeAlso: `WEBDOCS/truncate.html
`,
	},
	//line sql.y: 5077
	`CREATE USER`: {
		ShortDescription: `define a new user`,
		//line sql.y: 5078
		Category: hPriv,
		//line sql.y: 5079
		Text: `CREATE USER [IF NOT EXISTS] <name> [ [WITH] PASSWORD <passwd> ]
`,
		//line sql.y: 5080
		SeeAlso: `DROP USER, SHOW USERS, WEBDOCS/create-user.html
`,
	},
	//line sql.y: 5109
	`CREATE ROLE`: {
		ShortDescription: `define a new role`,
		//line sql.y: 5110
		Category: hPriv,
		//line sql.y: 5111
		Text: `CREATE ROLE [IF NOT EXISTS] <name>
`,
		//line sql.y: 5112
		SeeAlso: `DROP ROLE, SHOW ROLES
`,
	},
	//line sql.y: 5130
	`CREATE VIEW`: {
		ShortDescription: `create a new view`,
		//line sql.y: 5131
		Category: hDDL,
		//line sql.y: 5132
		Text: `CREATE [TEMPORARY | TEMP] VIEW <viewname> [( <colnames...> )] AS <source>
`,
		//line sql.y: 5133
		SeeAlso: `CREATE TABLE, SHOW CREATE, WEBDOCS/create-view.html
`,
	},
	//line sql.y: 5168
	`CREATE INDEX`: {
		ShortDescription: `create a new index`,
		//line sql.y: 5169
		Category: hDDL,
		//line sql.y: 5170
		Text: `
CREATE [UNIQUE | INVERTED] INDEX [IF NOT EXISTS] [<idxname>]
       ON <tablename> ( <colname> [ASC | DESC] [, ...] )
       [STORING ( <colnames...> )] [<interleave>]

Interleave clause:
   INTERLEAVE IN PARENT <tablename> ( <colnames...> ) [CASCADE | RESTRICT]

`,
		//line sql.y: 5178
		SeeAlso: `CREATE TABLE, SHOW INDEXES, SHOW CREATE,
WEBDOCS/create-index.html
`,
	},
	//line sql.y: 5407
	`RELEASE`: {
		ShortDescription: `complete a retryable block`,
		//line sql.y: 5408
		Category: hTxn,
		//line sql.y: 5409
		Text: `RELEASE [SAVEPOINT] cockroach_restart
`,
		//line sql.y: 5410
		SeeAlso: `SAVEPOINT, WEBDOCS/savepoint.html
`,
	},
	//line sql.y: 5418
	`RESUME JOBS`: {
		ShortDescription: `resume background jobs`,
		//line sql.y: 5419
		Category: hMisc,
		//line sql.y: 5420
		Text: `
RESUME JOBS <selectclause>
RESUME JOB <jobid>
`,
		//line sql.y: 5423
		SeeAlso: `SHOW JOBS, CANCEL JOBS, PAUSE JOBS
`,
	},
	//line sql.y: 5440
	`SAVEPOINT`: {
		ShortDescription: `start a retryable block`,
		//line sql.y: 5441
		Category: hTxn,
		//line sql.y: 5442
		Text: `SAVEPOINT cockroach_restart
`,
		//line sql.y: 5443
		SeeAlso: `RELEASE, WEBDOCS/savepoint.html
`,
	},
	//line sql.y: 5458
	`BEGIN`: {
		ShortDescription: `start a transaction`,
		//line sql.y: 5459
		Category: hTxn,
		//line sql.y: 5460
		Text: `
BEGIN [TRANSACTION] [ <txnparameter> [[,] ...] ]
START TRANSACTION [ <txnparameter> [[,] ...] ]

Transaction parameters:
   ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }
   PRIORITY { LOW | NORMAL | HIGH }

`,
		//line sql.y: 5468
		SeeAlso: `COMMIT, ROLLBACK, WEBDOCS/begin-transaction.html
`,
	},
	//line sql.y: 5481
	`COMMIT`: {
		ShortDescription: `commit the current transaction`,
		//line sql.y: 5482
		Category: hTxn,
		//line sql.y: 5483
		Text: `
COMMIT [TRANSACTION]
END [TRANSACTION]
`,
		//line sql.y: 5486
		SeeAlso: `BEGIN, ROLLBACK, WEBDOCS/commit-transaction.html
`,
	},
	//line sql.y: 5510
	`ROLLBACK`: {
		ShortDescription: `abort the current transaction`,
		//line sql.y: 5511
		Category: hTxn,
		//line sql.y: 5512
		Text: `ROLLBACK [TRANSACTION] [TO [SAVEPOINT] cockroach_restart]
`,
		//line sql.y: 5513
		SeeAlso: `BEGIN, COMMIT, SAVEPOINT, WEBDOCS/rollback-transaction.html
`,
	},
	//line sql.y: 5631
	`CREATE DATABASE`: {
		ShortDescription: `create a new database`,
		//line sql.y: 5632
		Category: hDDL,
		//line sql.y: 5633
		Text: `CREATE DATABASE [IF NOT EXISTS] <name>
`,
		//line sql.y: 5634
		SeeAlso: `WEBDOCS/create-database.html
`,
	},
	//line sql.y: 5703
	`INSERT`: {
		ShortDescription: `create new rows in a table`,
		//line sql.y: 5704
		Category: hDML,
		//line sql.y: 5705
		Text: `
INSERT INTO <tablename> [[AS] <name>] [( <colnames...> )]
       <selectclause>
       [ON CONFLICT [( <colnames...> )] {DO UPDATE SET ... [WHERE <expr>] | DO NOTHING}]
       [RETURNING <exprs...>]
`,
		//line sql.y: 5710
		SeeAlso: `UPSERT, UPDATE, DELETE, WEBDOCS/insert.html
`,
	},
	//line sql.y: 5729
	`UPSERT`: {
		ShortDescription: `create or replace rows in a table`,
		//line sql.y: 5730
		Category: hDML,
		//line sql.y: 5731
		Text: `
UPSERT INTO <tablename> [AS <name>] [( <colnames...> )]
       <selectclause>
       [RETURNING <exprs...>]
`,
		//line sql.y: 5735
		SeeAlso: `INSERT, UPDATE, DELETE, WEBDOCS/upsert.html
`,
	},
	//line sql.y: 5846
	`UPDATE`: {
		ShortDescription: `update rows of a table`,
		//line sql.y: 5847
		Category: hDML,
		//line sql.y: 5848
		Text: `
UPDATE <tablename> [[AS] <name>]
       SET ...
       [WHERE <expr>]
       [ORDER BY <exprs...>]
       [LIMIT <expr>]
       [RETURNING <exprs...>]
`,
		//line sql.y: 5855
		SeeAlso: `INSERT, UPSERT, DELETE, WEBDOCS/update.html
`,
	},
	//line sql.y: 6080
	`<SELECTCLAUSE>`: {
		ShortDescription: `access tabular data`,
		//line sql.y: 6081
		Category: hDML,
		//line sql.y: 6082
		Text: `
Select clause:
  TABLE <tablename>
  VALUES ( <exprs...> ) [ , ... ]
  SELECT ... [ { INTERSECT | UNION | EXCEPT } [ ALL | DISTINCT ] <selectclause> ]
`,
	},
	//line sql.y: 6093
	`SELECT`: {
		ShortDescription: `retrieve rows from a data source and compute a result`,
		//line sql.y: 6094
		Category: hDML,
		//line sql.y: 6095
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
		//line sql.y: 6107
		SeeAlso: `WEBDOCS/select-clause.html
`,
	},
	//line sql.y: 6182
	`TABLE`: {
		ShortDescription: `select an entire table`,
		//line sql.y: 6183
		Category: hDML,
		//line sql.y: 6184
		Text: `TABLE <tablename>
`,
		//line sql.y: 6185
		SeeAlso: `SELECT, VALUES, WEBDOCS/table-expressions.html
`,
	},
	//line sql.y: 6497
	`VALUES`: {
		ShortDescription: `select a given set of values`,
		//line sql.y: 6498
		Category: hDML,
		//line sql.y: 6499
		Text: `VALUES ( <exprs...> ) [, ...]
`,
		//line sql.y: 6500
		SeeAlso: `SELECT, TABLE, WEBDOCS/table-expressions.html
`,
	},
	//line sql.y: 6609
	`<SOURCE>`: {
		ShortDescription: `define a data source for SELECT`,
		//line sql.y: 6610
		Category: hDML,
		//line sql.y: 6611
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
		//line sql.y: 6633
		SeeAlso: `WEBDOCS/table-expressions.html
`,
	},
}
