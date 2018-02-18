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
$dbh->do("SET FOREIGN_KEY_CHECKS=0;");

for (my $i=0; $i< 40; $i++) {
    my $dbname = sprintf("testdb_%04d", $i);
    my $query = "CREATE SCHEMA IF NOT EXISTS $dbname";
    print "$query\n";
    $dbh->do($query);
    for (my $j=0; $j < 100; $j++) {
        my $table = sprintf("table_%04d", $j);
        $query = "CREATE TABLE IF NOT EXISTS $dbname.$table (id int, f2 varchar(25))";
        print "$query\n";
        $dbh->do($query);
    }
}

$dbh->do("SET FOREIGN_KEY_CHECKS=1;");
