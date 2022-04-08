%%Estefanía Pérez Hidalgo
%Subset exercise.
%
%Possible query: 'substring([1,2,3,4,5,6,7],[4,5,6]).'.


substring(List, List).
substring(List, Rest) :-
   empty(List, Rest).

empty([_|Tail], Rest) :-
   substring(Tail, Rest).
empty([Head|Tail], [Head|Rest]) :-
   empty(Tail, Rest).
