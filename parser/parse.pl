#!/usr/bin/env perl

use strict;
use warnings;
use autodie;
use feature qw/say/;
use HTML::TableExtract;
use Text::CSV qw/csv/;

open my $fh, '<', "company.html";

my $html_string = do {
    local $/;
    <$fh>;
};

my @headers = (qw(company employees revenue headquarters founded), "market cap");
my @data = [q/country/, @headers];

my $te = HTML::TableExtract->new(
    headers => \@headers 
);
$te->parse($html_string);

foreach my $row ($te->rows) {
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
foreach my $row ($te->rows) {
    if ($row->[0] =~ m/title="([^"]+)"/){
        unshift @{$data[$i]}, $1;
    }
    $i++;
}

csv (in => \@data, out => "company.csv", sep_char=> ";");


