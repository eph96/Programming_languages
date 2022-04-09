%%Estefanía Pérez Hidalgo
%Restaurat exercise.
%
%Possible query: 'maxCalories(5000,paella,bacalao,naranja).'.

appetizers(paella).
appetizers(gazpacho).
appetizers(consume).
meat(fileteCerdo).
meat(polloAsado).
fish(trucha).
fish(bacalao).
dessert(flan).
dessert(nuecesMiel).
dessert(naranja).
calories(paella,350).
calories(gazpacho,175).
calories(consume,300).
calories(fileteCerdo,200).
calories(polloAsado,150).
calories(trucha,108).
calories(bacalao,100).
calories(flan,150).
calories(nuecesMiel,200).
calories(naranja,173).


main_course(X):-
    meat(X).

main_course(X):-
    fish(X).

complete_meal(X,Y,Z):-
    appetizers(X),
    main_course(Y),
    dessert(Z).

caloriesAmount(Food1,Food2,Food3,Value):-
    calories(Food1,Calories1),
    calories(Food2,Calories2),
    calories(Food3,Calories3),
    Value is Calories1+Calories2+Calories3,
    complete_meal(Food1,Food2,Food3).

maxCalories(Max,Food1,Food2,Food3):-
    caloriesAmount(Food1,Food2,Food3,Value),
    Value<Max.
