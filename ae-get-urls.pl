#!/usr/bin/perl
use strict;
use AnyEvent;
use AnyEvent::HTTP;
use Time::HiRes qw/time/;
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

while(my $url = <>) {
	chomp($url);
	if($url !~ /($RE{URI}{HTTP})/) {
		warn "URL: \"$url\" is not valid\n";
	} else {
		push @urls, $url;
	}
}

#warn "URL: $_\n" for (@urls);

my $start_time = time;
my $timeout    = 2;
my $count      = 0;
my %Count;
my $result;

my $cv = AnyEvent->condvar();
$cv->begin(sub { shift->send($result) });

for my $url (@urls) {

	$cv->begin();
	$Count{$url} = ++$count;
	my $now = time;

	my $request;
	$request = http_request(
		GET     => $url,
		timeout => 2,      # seconds
		sub {
			my ($body, $hdr) = @_;

			#$count++;
			if($hdr->{Status} =~ /^2/) {
				push(@$result, join("\t", (qq/[$Count{$url}] "$url"/,
							" length: ",
							$hdr->{'content-length'},
							" time: ",
							sprintf("%.3f", time - $now),
							"ms"))
				);

				#print $body;
			} else {
				push(@$result, join("", qq/[$Count{$url}] "$url"/,
						": (",
						$hdr->{Status},
						") ",
						$hdr->{Reason})
				);
			}
			$cv->end;
			undef $request;
			}
	);
}

$cv->end;
my $res = $cv->recv;

print "\n" . "-" x 10 . "\n" . join("\n", @$res) . "\n" . "-" x 10 . "\n" if(defined($res));
print "Total elapsed time: " . sprintf("%.3f", time - $start_time) . "ms\n";

__END__
