#!/usr/bin/perl
#
use Data::Dumper;
use DBI;

my $database="";
my $hostname="127.0.0.1";
my $port=3308;
my $user="root";
my $password="";

my $dsn = "DBI:mysql:database=$database;host=$hostname;port=$port";
my $dbh = DBI->connect($dsn, $user, $password);

for (my $i=0; $i< 40; $i++) {
    my $dbname = sprintf("testdb_%04d", $i);
    $dbh->do("CREATE SCHEMA IF NOT EXISTS $dbname");
    for (my $j=0; $j < 100; $j++) {
        my $table = sprintf("table_%04d", $j);
        $dbh->do("CREATE TABLE IF NOT EXISTS $dbname.$table (id int, f2 varchar(25))");
    }
}

