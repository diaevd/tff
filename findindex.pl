#!/usr/bin/perl

use strict;
use EvgenyDuzhakovFindIndex;

my @arr = qw(-30 -20 -10 1 10 20 30 40 50 60 70 80 90 100);
print "\@arr: " . join(", ", map { "[$_] => $arr[$_]" } (0 .. $#arr)) . "\n";

my $finder = EvgenyDuzhakovFindIndex->new();
my ($index, $passes) = $finder->find(-5, \@arr);
print "index=$index passes=$passes\n";

1;

__END__
