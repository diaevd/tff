#!/usr/bin/perl
use strict;
use Regexp::Common qw/URI/;
use File::Basename;

if(-t STDIN) {
	my $self_name = basename($0);
	die <<EOT;
Get urls by list from <STDIN>

Usage:
\tcat <urls.txt> | $self_name
or
\techo <url> | $self_name
EOT
}

my @urls = ();

while (my $url = <>) {
	chomp($url);
	if($url !~ /($RE{URI}{HTTP})/) {
		warn "INVALID URL: \"$url\n";
	} else {
		push @urls, $url;
	}
}

warn "URL: $_\n" for (@urls);

1;
__END__
