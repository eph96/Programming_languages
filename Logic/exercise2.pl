%%Estefanía Pérez Hidalgo
%Even and odd lists exercise.
%Possible query: 'largopar([1,2,3,4]).', or 'largoimpar([1,2,3,4,5]).'.

largopar([]):-
    write('Largo par').

largopar([_|Tail]):-
    largoimpar(Tail).

largoimpar([_]):-
    write('Largo impar').

largoimpar([_|Tail]):-
    largopar(Tail).
