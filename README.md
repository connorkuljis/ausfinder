# Importing ASIC Business Names Dataset into SQLite Guide

## Get the dataset

[ASIC Business Names Dataset](https://www.data.gov.au/dataset/ds-dga-bc515135-4bb6-4d50-957a-3713709a76d3/)

* Note: the dataset is tab separated.

## Setup the SQLite database

Create a SQLite database and start the command line shell.

`sqlite3 db.sqlite3`
  
## Import the CSV file into a new table.

Now that we are in the shell, to import the dataset first we need to enable tab mode. Then we can create the table using the `.import` command. It will take a while to import the file so wait a little and verify the table creation by viewing the schema.

* Note: omit semi-colons on sqlite commands.

```
.mode tabs
.import {{ filename }} business_names
```

View the schema

`.schema business_names`


Verfify table entries using count

`select count(*) from business_names;`

## Notes

### Line Mode

```
.mode line
select trim(BN_NAME) AS 'Business', BN_ABN as ABN from business LIMIT 10;`
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

```
sqlite> select distinct BN_STATUS from business;
Registered
Deregistered
```
