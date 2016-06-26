## no critic (RCS,prototypes)

package EvgenyDuzhakovFindIndex;

use 5.006000;

use strict;
use warnings;

our $VERSION = '0.01';

sub new {
	my $proto = shift;
	my $class = ref($proto) || $proto;
	my $self = {};
	$self->{DEBUG} = 0;
	bless($self, $class);
}

sub Debug {
	my $self = shift;
	$self->{DEBUG} = $_[0] if(@_);
	return $self->{DEBUG}
}

sub find ($\@) {
	my $self = shift;
	my ($needle, $aref) = @_;

	my ($index, $low, $high, $passes) = (0, 0, scalar(@$aref), 1);

	while($low < $high) {
		my $cur = int(($high - $low) / 2 + $low);
		if($needle > $aref->[$cur]) {
			$low = $cur + 1;
		} else {
			$high = $cur;
		}
		$passes++;
	}

	$high = scalar(@$aref);
	if($low == 0) { $index = 0; }
	elsif($low == $high) { $index = $high - 1; }
	elsif($aref->[$low] == $needle) { $index = $low; }
	elsif(abs($needle - $aref->[$low - 1]) <= abs($aref->[$low] - $needle)) { $index = $low - 1; }
	else { $index = $low; };

	print scalar((caller(0))[3]) . "(" . $needle . "): index=$index, passes=$passes, aref[$index]=$aref->[$index]" . "\n" if($self->Debug());

	return ($index, $passes);
}

1;

__END__

=head1 NAME

=head1 SYNOPSIS

=head1 DESCRIPTION

=head1 EXPORT

=head1 SUBROUTINES/METHODS


=head1 CONFIGURATION AND ENVIRONMENT

=head1 DEPENDENCIES

Perl 5.8.

=head1 INCOMPATIBILITIES

=head1 AUTHOR

Evgeny Duzhakov, C<< <diaevd at gmail.com> >>

If the documentation fails to answer your question, or if you have a comment
or suggestion, send me an email.


=head1 DIAGNOSTICS


=head1 BUGS AND LIMITATIONS

=head1 SUPPORT


=head1 ACKNOWLEDGEMENTS

=head1 LICENSE AND COPYRIGHT

Copyright 2016 Evgeny Duzhakov.

This program is free software; you can redistribute it and/or modify it
under the terms of either: the GNU General Public License as published
by the Free Software Foundation; or the Artistic License.

See http://dev.perl.org/licenses/ for more information.

=cut
