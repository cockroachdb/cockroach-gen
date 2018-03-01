// Code generated by parser/all_keywords.awk. DO NOT EDIT.
// GENERATED FILE DO NOT EDIT

package lex

var Keywords = map[string]struct {
	Tok int
	Cat string
}{
	"abort":                     {ABORT, "U"},
	"action":                    {ACTION, "U"},
	"add":                       {ADD, "U"},
	"admin":                     {ADMIN, "U"},
	"all":                       {ALL, "R"},
	"alter":                     {ALTER, "U"},
	"analyse":                   {ANALYSE, "R"},
	"analyze":                   {ANALYZE, "R"},
	"and":                       {AND, "R"},
	"annotate_type":             {ANNOTATE_TYPE, "C"},
	"any":                       {ANY, "R"},
	"array":                     {ARRAY, "R"},
	"as":                        {AS, "R"},
	"asc":                       {ASC, "R"},
	"asymmetric":                {ASYMMETRIC, "R"},
	"at":                        {AT, "U"},
	"backup":                    {BACKUP, "U"},
	"begin":                     {BEGIN, "U"},
	"between":                   {BETWEEN, "C"},
	"bigint":                    {BIGINT, "C"},
	"bigserial":                 {BIGSERIAL, "C"},
	"bit":                       {BIT, "C"},
	"blob":                      {BLOB, "U"},
	"bool":                      {BOOL, "C"},
	"boolean":                   {BOOLEAN, "C"},
	"both":                      {BOTH, "R"},
	"by":                        {BY, "U"},
	"bytea":                     {BYTEA, "C"},
	"bytes":                     {BYTES, "C"},
	"cache":                     {CACHE, "U"},
	"cancel":                    {CANCEL, "U"},
	"cascade":                   {CASCADE, "U"},
	"case":                      {CASE, "R"},
	"cast":                      {CAST, "R"},
	"char":                      {CHAR, "C"},
	"character":                 {CHARACTER, "C"},
	"characteristics":           {CHARACTERISTICS, "C"},
	"check":                     {CHECK, "R"},
	"cluster":                   {CLUSTER, "U"},
	"coalesce":                  {COALESCE, "C"},
	"collate":                   {COLLATE, "R"},
	"collation":                 {COLLATION, "T"},
	"column":                    {COLUMN, "R"},
	"columns":                   {COLUMNS, "U"},
	"comment":                   {COMMENT, "U"},
	"commit":                    {COMMIT, "U"},
	"committed":                 {COMMITTED, "U"},
	"compact":                   {COMPACT, "U"},
	"configuration":             {CONFIGURATION, "U"},
	"configurations":            {CONFIGURATIONS, "U"},
	"configure":                 {CONFIGURE, "U"},
	"conflict":                  {CONFLICT, "U"},
	"constraint":                {CONSTRAINT, "R"},
	"constraints":               {CONSTRAINTS, "U"},
	"copy":                      {COPY, "U"},
	"covering":                  {COVERING, "U"},
	"create":                    {CREATE, "R"},
	"cross":                     {CROSS, "T"},
	"csv":                       {CSV, "U"},
	"cube":                      {CUBE, "U"},
	"current":                   {CURRENT, "U"},
	"current_catalog":           {CURRENT_CATALOG, "R"},
	"current_date":              {CURRENT_DATE, "R"},
	"current_role":              {CURRENT_ROLE, "R"},
	"current_schema":            {CURRENT_SCHEMA, "R"},
	"current_time":              {CURRENT_TIME, "R"},
	"current_timestamp":         {CURRENT_TIMESTAMP, "R"},
	"current_user":              {CURRENT_USER, "R"},
	"cycle":                     {CYCLE, "U"},
	"data":                      {DATA, "U"},
	"database":                  {DATABASE, "U"},
	"databases":                 {DATABASES, "U"},
	"date":                      {DATE, "C"},
	"day":                       {DAY, "U"},
	"deallocate":                {DEALLOCATE, "U"},
	"dec":                       {DEC, "C"},
	"decimal":                   {DECIMAL, "C"},
	"default":                   {DEFAULT, "R"},
	"deferrable":                {DEFERRABLE, "R"},
	"delete":                    {DELETE, "U"},
	"desc":                      {DESC, "R"},
	"discard":                   {DISCARD, "U"},
	"distinct":                  {DISTINCT, "R"},
	"do":                        {DO, "R"},
	"double":                    {DOUBLE, "U"},
	"drop":                      {DROP, "U"},
	"else":                      {ELSE, "R"},
	"encoding":                  {ENCODING, "U"},
	"end":                       {END, "R"},
	"except":                    {EXCEPT, "R"},
	"execute":                   {EXECUTE, "U"},
	"exists":                    {EXISTS, "C"},
	"experimental":              {EXPERIMENTAL, "U"},
	"experimental_audit":        {EXPERIMENTAL_AUDIT, "U"},
	"experimental_fingerprints": {EXPERIMENTAL_FINGERPRINTS, "U"},
	"experimental_replica":      {EXPERIMENTAL_REPLICA, "U"},
	"explain":                   {EXPLAIN, "U"},
	"extract":                   {EXTRACT, "C"},
	"extract_duration":          {EXTRACT_DURATION, "C"},
	"false":                     {FALSE, "R"},
	"family":                    {FAMILY, "T"},
	"fetch":                     {FETCH, "R"},
	"filter":                    {FILTER, "U"},
	"first":                     {FIRST, "U"},
	"float":                     {FLOAT, "C"},
	"float4":                    {FLOAT4, "C"},
	"float8":                    {FLOAT8, "C"},
	"following":                 {FOLLOWING, "U"},
	"for":                       {FOR, "R"},
	"force_index":               {FORCE_INDEX, "U"},
	"foreign":                   {FOREIGN, "R"},
	"from":                      {FROM, "R"},
	"full":                      {FULL, "T"},
	"gin":                       {GIN, "U"},
	"grant":                     {GRANT, "R"},
	"grants":                    {GRANTS, "U"},
	"greatest":                  {GREATEST, "C"},
	"group":                     {GROUP, "R"},
	"grouping":                  {GROUPING, "C"},
	"having":                    {HAVING, "R"},
	"high":                      {HIGH, "U"},
	"histogram":                 {HISTOGRAM, "U"},
	"hour":                      {HOUR, "U"},
	"if":                        {IF, "C"},
	"ifnull":                    {IFNULL, "C"},
	"ilike":                     {ILIKE, "T"},
	"import":                    {IMPORT, "U"},
	"in":                        {IN, "R"},
	"increment":                 {INCREMENT, "U"},
	"incremental":               {INCREMENTAL, "U"},
	"index":                     {INDEX, "R"},
	"indexes":                   {INDEXES, "U"},
	"inet":                      {INET, "C"},
	"initially":                 {INITIALLY, "R"},
	"inner":                     {INNER, "T"},
	"insert":                    {INSERT, "U"},
	"int":                       {INT, "C"},
	"int2":                      {INT2, "C"},
	"int2vector":                {INT2VECTOR, "U"},
	"int4":                      {INT4, "C"},
	"int64":                     {INT64, "C"},
	"int8":                      {INT8, "C"},
	"integer":                   {INTEGER, "C"},
	"interleave":                {INTERLEAVE, "U"},
	"intersect":                 {INTERSECT, "R"},
	"interval":                  {INTERVAL, "C"},
	"into":                      {INTO, "R"},
	"inverted":                  {INVERTED, "U"},
	"is":                        {IS, "T"},
	"isolation":                 {ISOLATION, "U"},
	"job":                       {JOB, "U"},
	"jobs":                      {JOBS, "U"},
	"join":                      {JOIN, "T"},
	"json":                      {JSON, "C"},
	"jsonb":                     {JSONB, "C"},
	"key":                       {KEY, "U"},
	"keys":                      {KEYS, "U"},
	"kv":                        {KV, "U"},
	"lateral":                   {LATERAL, "R"},
	"lc_collate":                {LC_COLLATE, "U"},
	"lc_ctype":                  {LC_CTYPE, "U"},
	"leading":                   {LEADING, "R"},
	"least":                     {LEAST, "C"},
	"left":                      {LEFT, "T"},
	"less":                      {LESS, "U"},
	"level":                     {LEVEL, "U"},
	"like":                      {LIKE, "T"},
	"limit":                     {LIMIT, "R"},
	"list":                      {LIST, "U"},
	"local":                     {LOCAL, "U"},
	"localtime":                 {LOCALTIME, "R"},
	"localtimestamp":            {LOCALTIMESTAMP, "R"},
	"low":                       {LOW, "U"},
	"match":                     {MATCH, "U"},
	"maxvalue":                  {MAXVALUE, "T"},
	"minute":                    {MINUTE, "U"},
	"minvalue":                  {MINVALUE, "T"},
	"month":                     {MONTH, "U"},
	"name":                      {NAME, "C"},
	"names":                     {NAMES, "U"},
	"nan":                       {NAN, "U"},
	"natural":                   {NATURAL, "T"},
	"next":                      {NEXT, "U"},
	"no":                        {NO, "U"},
	"no_index_join":             {NO_INDEX_JOIN, "U"},
	"normal":                    {NORMAL, "U"},
	"not":                       {NOT, "R"},
	"nothing":                   {NOTHING, "R"},
	"null":                      {NULL, "R"},
	"nullif":                    {NULLIF, "C"},
	"nulls":                     {NULLS, "U"},
	"numeric":                   {NUMERIC, "C"},
	"of":                        {OF, "U"},
	"off":                       {OFF, "U"},
	"offset":                    {OFFSET, "R"},
	"oid":                       {OID, "U"},
	"on":                        {ON, "R"},
	"only":                      {ONLY, "R"},
	"option":                    {OPTION, "U"},
	"options":                   {OPTIONS, "U"},
	"or":                        {OR, "R"},
	"order":                     {ORDER, "R"},
	"ordinality":                {ORDINALITY, "U"},
	"out":                       {OUT, "C"},
	"outer":                     {OUTER, "T"},
	"over":                      {OVER, "U"},
	"overlaps":                  {OVERLAPS, "T"},
	"overlay":                   {OVERLAY, "C"},
	"owned":                     {OWNED, "U"},
	"parent":                    {PARENT, "U"},
	"partial":                   {PARTIAL, "U"},
	"partition":                 {PARTITION, "U"},
	"password":                  {PASSWORD, "U"},
	"pause":                     {PAUSE, "U"},
	"physical":                  {PHYSICAL, "U"},
	"placing":                   {PLACING, "R"},
	"plans":                     {PLANS, "U"},
	"position":                  {POSITION, "C"},
	"preceding":                 {PRECEDING, "U"},
	"precision":                 {PRECISION, "C"},
	"prepare":                   {PREPARE, "U"},
	"primary":                   {PRIMARY, "R"},
	"priority":                  {PRIORITY, "U"},
	"queries":                   {QUERIES, "U"},
	"query":                     {QUERY, "U"},
	"range":                     {RANGE, "U"},
	"read":                      {READ, "U"},
	"real":                      {REAL, "C"},
	"recursive":                 {RECURSIVE, "U"},
	"ref":                       {REF, "U"},
	"references":                {REFERENCES, "R"},
	"regclass":                  {REGCLASS, "U"},
	"regnamespace":              {REGNAMESPACE, "U"},
	"regproc":                   {REGPROC, "U"},
	"regprocedure":              {REGPROCEDURE, "U"},
	"regtype":                   {REGTYPE, "U"},
	"release":                   {RELEASE, "U"},
	"rename":                    {RENAME, "U"},
	"repeatable":                {REPEATABLE, "U"},
	"reset":                     {RESET, "U"},
	"restore":                   {RESTORE, "U"},
	"restrict":                  {RESTRICT, "U"},
	"resume":                    {RESUME, "U"},
	"returning":                 {RETURNING, "R"},
	"revoke":                    {REVOKE, "U"},
	"right":                     {RIGHT, "T"},
	"role":                      {ROLE, "R"},
	"roles":                     {ROLES, "U"},
	"rollback":                  {ROLLBACK, "U"},
	"rollup":                    {ROLLUP, "U"},
	"row":                       {ROW, "C"},
	"rows":                      {ROWS, "U"},
	"savepoint":                 {SAVEPOINT, "U"},
	"scatter":                   {SCATTER, "U"},
	"schema":                    {SCHEMA, "U"},
	"schemas":                   {SCHEMAS, "U"},
	"scrub":                     {SCRUB, "U"},
	"search":                    {SEARCH, "U"},
	"second":                    {SECOND, "U"},
	"select":                    {SELECT, "R"},
	"sequence":                  {SEQUENCE, "U"},
	"sequences":                 {SEQUENCES, "U"},
	"serial":                    {SERIAL, "C"},
	"serial2":                   {SERIAL2, "C"},
	"serial4":                   {SERIAL4, "C"},
	"serial8":                   {SERIAL8, "C"},
	"serializable":              {SERIALIZABLE, "U"},
	"session":                   {SESSION, "U"},
	"session_user":              {SESSION_USER, "R"},
	"sessions":                  {SESSIONS, "U"},
	"set":                       {SET, "U"},
	"setting":                   {SETTING, "U"},
	"settings":                  {SETTINGS, "U"},
	"show":                      {SHOW, "U"},
	"similar":                   {SIMILAR, "T"},
	"simple":                    {SIMPLE, "U"},
	"smallint":                  {SMALLINT, "C"},
	"smallserial":               {SMALLSERIAL, "C"},
	"snapshot":                  {SNAPSHOT, "U"},
	"some":                      {SOME, "R"},
	"split":                     {SPLIT, "U"},
	"sql":                       {SQL, "U"},
	"start":                     {START, "U"},
	"statistics":                {STATISTICS, "U"},
	"status":                    {STATUS, "U"},
	"stdin":                     {STDIN, "U"},
	"store":                     {STORE, "U"},
	"stored":                    {STORED, "R"},
	"storing":                   {STORING, "U"},
	"strict":                    {STRICT, "U"},
	"string":                    {STRING, "C"},
	"substring":                 {SUBSTRING, "C"},
	"symmetric":                 {SYMMETRIC, "R"},
	"syntax":                    {SYNTAX, "U"},
	"system":                    {SYSTEM, "U"},
	"table":                     {TABLE, "R"},
	"tables":                    {TABLES, "U"},
	"temp":                      {TEMP, "U"},
	"template":                  {TEMPLATE, "U"},
	"temporary":                 {TEMPORARY, "U"},
	"testing_ranges":            {TESTING_RANGES, "U"},
	"testing_relocate":          {TESTING_RELOCATE, "U"},
	"text":                      {TEXT, "U"},
	"than":                      {THAN, "U"},
	"then":                      {THEN, "R"},
	"time":                      {TIME, "C"},
	"timestamp":                 {TIMESTAMP, "C"},
	"timestamptz":               {TIMESTAMPTZ, "C"},
	"to":                        {TO, "R"},
	"trace":                     {TRACE, "U"},
	"trailing":                  {TRAILING, "R"},
	"transaction":               {TRANSACTION, "U"},
	"treat":                     {TREAT, "C"},
	"trim":                      {TRIM, "C"},
	"true":                      {TRUE, "R"},
	"truncate":                  {TRUNCATE, "U"},
	"type":                      {TYPE, "U"},
	"unbounded":                 {UNBOUNDED, "U"},
	"uncommitted":               {UNCOMMITTED, "U"},
	"union":                     {UNION, "R"},
	"unique":                    {UNIQUE, "R"},
	"unknown":                   {UNKNOWN, "U"},
	"update":                    {UPDATE, "U"},
	"upsert":                    {UPSERT, "U"},
	"use":                       {USE, "U"},
	"user":                      {USER, "R"},
	"users":                     {USERS, "U"},
	"using":                     {USING, "R"},
	"uuid":                      {UUID, "C"},
	"valid":                     {VALID, "U"},
	"validate":                  {VALIDATE, "U"},
	"value":                     {VALUE, "U"},
	"values":                    {VALUES, "C"},
	"varchar":                   {VARCHAR, "C"},
	"variadic":                  {VARIADIC, "R"},
	"varying":                   {VARYING, "U"},
	"view":                      {VIEW, "R"},
	"virtual":                   {VIRTUAL, "R"},
	"when":                      {WHEN, "R"},
	"where":                     {WHERE, "R"},
	"window":                    {WINDOW, "R"},
	"with":                      {WITH, "R"},
	"within":                    {WITHIN, "U"},
	"without":                   {WITHOUT, "U"},
	"work":                      {WORK, "R"},
	"write":                     {WRITE, "U"},
	"year":                      {YEAR, "U"},
	"zone":                      {ZONE, "U"},
}
