#!/usr/bin/env perl

use strict;
use warnings;
use autodie;
use feature qw/say/;
use HTML::TableExtract;
use Text::CSV qw/csv/;
use LWP::Simple qw/get/;
my $DATA_URL = q[https://en.wikipedia.org/wiki/List_of_largest_Internet_companies];

my $html_string = get($DATA_URL);
die "Can't get $DATA_URL" unless $html_string;

my @headers = (qw(company employees revenue headquarters founded));
my @data = [q/country/, @headers, q/marketcap/];
@headers = (@headers, "market cap");

my $te = HTML::TableExtract->new(
    headers => \@headers 
);
$te->parse($html_string);

foreach my $row ($te->rows) {
   no warnings 'uninitialized';
   @$row = map {s/^\s*//; s/\s*$//; $_} @$row;
   $row->[1] =~ s/,//;
   $row->[2] =~ s/[^.\d]//g;
   $row->[5] =~ s/[^.\d]//g;
   push @data, $row;
}

$te = HTML::TableExtract->new(
    headers => ['company'], 
    keep_html => 1,
);
$te->parse($html_string);

my $i = 1;
my $ts = ($te->tables)[1];
foreach my $row ($ts->rows) {
    if ($row->[0] =~ m/title="([^"]+)"/){
        unshift @{$data[$i]}, $1;
    }
    $i++;
}

csv (in => \@data, out => "company.csv", sep_char=> ";");


