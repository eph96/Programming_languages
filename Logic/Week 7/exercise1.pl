%%Estefanía Pérez Hidalgo
%Aunt exercise.
%
%Possible query: 'aunt(Aunt,Niece).'.

femenino(petra).
femenino(carmen).
femenino(maria).
femenino(rosa).
femenino(ana).
femenino(cande).
madre(petra, juan).
madre(petra, jose).
madre(petra, rosa).
madre(carmen, pedro).
madre(maria, ana).
madre(maria, enrique).
madre(rosa, raul).
madre(rosa, alfonso).
madre(rosa, cande).
masculino(angel).
masculino(pepe).
masculino(juan).
masculino(pedro).
masculino(enrique).
masculino(raul).
masculino(alfonso).
padre(angel, juan).
padre(angel, rosa).
padre(pepe, pedro).
padre(juan, ana).
padre(juan, enrique).
padre(pedro, raul).
padre(pedro, alfonso).
padre(pedro, cande).

% Rosa is Enrique's aunt. Enrique's father is Juan and his mother is
% Marina. Juan and Rosa are siblings.

aunt(X,S) :-
    padre(Z,S),
    padre(H,Z),
    padre(H,X),
    Z\=X,
    femenino(X).

