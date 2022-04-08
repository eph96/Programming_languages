%%Estefanía Pérez Hidalgo
%Sum exercise.
%Possible query: 'sumlist([1,2,3],Result).'

sumlist([],0).
sumlist([Head|Tail],Result):-
		sum(Tail,Result2),
		Result is Result2+Head.

