# Importing ASIC Business Names Dataset into SQLite Guide

## Get the dataset

[ASIC Business Names Dataset](https://www.data.gov.au/dataset/ds-dga-bc515135-4bb6-4d50-957a-3713709a76d3/)

* Note: the dataset is tab separated.

## Setup the SQLite database

Create a SQLite database and start the command line shell.

`sqlite3 data/db.sqlite3`
  
## Import the CSV file into a new table.

Now that we are in the shell, first we need to enable tab mode. Then we can create the table using the `.import` command. It will take a while to import the file, so wait a little while and then verify the table creation by viewing the schema.

* Note: omit semi-colons on sqlite commands.

```
.mode tabs
.import data/BUSINESS_NAMES_202501.csv business_names
```

## Verfify table creation and population

Verify the table creation by viewing the created schema.

`.schema business_names`

Check results

```sql
select count(*) from business_names;
```

## Add Full Text Search

[SQLite FTS5 Extension ](https://www.sqlite.org/fts5.html)

Create virtual table with fts5.

```sql
CREATE VIRTUAL TABLE businesses USING fts5(business_name, abn);
```

Populate the virtual table using insert select.

```sql
INSERT INTO
     businesses
     SELECT
          trim(BN_NAME), trim(BN_ABN)
     FROM
          business_names;
```


## Notes

### Line Mode

`.mode line`

```sql
select trim(BN_NAME) as 'Business', BN_ABN as 'ABN' from business_names limit 10;
```

Outputs

```
Business =
     ABN = 87387704324

Business = SILENT SCISSORZ
     ABN = 76643277300

Business = LITTLE MIRACLES PRESCHOOL & LONG DAY CARE (POINT CLARE, SWANSEA, MT RIVERVIEW & WAMBERAL)
     ABN = 23979823212

Business = A Cut Above Painting & Texture Coating
     ABN = 86634681397

Business = HOMSAFE
     ABN = 56098948915

Business = COASTAL  EARTH WORKS
     ABN = 88573118334

Business = Easy Settlements
     ABN = 47613497233

Business = Nourishing Your World
     ABN =
```

I find `.mode line` easy to read for a few reasons, for example it mitigate line wrapping and every entry is delimited by a new line.

From this example there are a few oddites:

* Emtpy Business Names
* Empty Australian Business Numbers

---

* use `trim()` to strip leading and trailing whitespace. This is especially useful for Business Names.


## Example Queries

Query `business_names` table

```sql
select count(*) from business_names;
```

```sql
select distinct BN_STATUS from business_names;
```

```sql
select distinct BN_STATE_OF_REG from business_names;
```

```sql
select trim(BN_NAME) as 'Business', BN_ABN as 'ABN' from business_names limit 10;
```

---

Query `business` virtual table with full text search.

```sql
select * from businesses WHERE business_name MATCH 'software';
```
