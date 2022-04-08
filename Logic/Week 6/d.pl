%%Estefanía Pérez Hidalgo
%Invert exercise.
%
%Possible query: 'beginning([6,5,4],[4,5,6]).'.


beginning(List1,List2):-
   comparison(List1,List2,[]).

comparison(_,List,List).
comparison(List, Rest,New) :-
   invert(List, Rest,New).

invert([Head|Tail], Rest,New) :-
   comparison(Tail, Rest,[Head|New]).
