#!/usr/bin/perl

use strict;
use EvgenyDuzhakovFindIndex;
use Test::More tests => 11;

my @arr = qw(-30 -20 -10 1 10 20 30 40 50 60 70 80 90 100);
print "\@arr: " . join(", ", map { "[$_] => $arr[$_]" } (0 .. $#arr)) . "\n";

my $finder = EvgenyDuzhakovFindIndex->new();
#$finder->Debug(1); # enable debug print
my ($index, $passes) = (0, 0);

($index, $passes) = $finder->find(-5, \@arr);
is($arr[$index], -10, "find '-5': [$index] => $arr[$index]), passes=$passes");

($index, $passes) = $finder->find(-3, \@arr);
is($arr[$index], 1, "find '-3': [$index] => $arr[$index]), passes=$passes");

($index, $passes) = $finder->find(-20, \@arr);
is($arr[$index], -20, "find '-20': [$index] => $arr[$index]), passes=$passes");

($index, $passes) = $finder->find(-21, \@arr);
is($arr[$index], -20, "find '-21': [$index] => $arr[$index]), passes=$passes");

($index, $passes) = $finder->find(20, \@arr);
is($arr[$index], 20, "find '20': [$index] => $arr[$index]), passes=$passes");

($index, $passes) = $finder->find(21, \@arr);
is($arr[$index], 20, "find '21': [$index] => $arr[$index]), passes=$passes");

($index, $passes) = $finder->find(25, \@arr);
is($arr[$index], 20, "find '25': [$index] => $arr[$index]), passes=$passes");

($index, $passes) = $finder->find(26, \@arr);
is($arr[$index], 30, "find '26': [$index] => $arr[$index]), passes=$passes");

($index, $passes) = $finder->find(29, \@arr);
is($arr[$index], 30, "find '29': [$index] => $arr[$index]), passes=$passes");

($index, $passes) = $finder->find(-35, \@arr);
is($arr[$index], -30, "find '-35': [$index] => $arr[$index]), passes=$passes");

($index, $passes) = $finder->find(110, \@arr);
is($arr[$index], 100, "find '110': [$index] => $arr[$index]), passes=$passes");

done_testing();

1;

__END__
