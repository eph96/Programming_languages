(ns exercise1.core)

(defn factorial [x]
  (if (= x 1)
    x
    (* x (factorial (- x 1)))
    )
  )
