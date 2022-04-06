%%Estefanía Pérez Hidalgo
%Sum exercise.
%Posible query: 'sum([1,2,3],Result).'



sum([],0).
sum([A|Tail],S):-
		sum(Tail,S2),
		S is S2+A.

