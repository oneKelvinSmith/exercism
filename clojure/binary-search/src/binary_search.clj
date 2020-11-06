(ns binary-search)

(defn spy [tag value] (prn tag value) value)

(defn middle [collection]
  (quot (count collection) 2))

(defn search-for
  [item collection]
  (loop [items (vec collection)
         start-index 0]
    (if (empty? items) (throw (Exception. "not found")))

    (let [middle-index (middle items)
          middle-item (get items middle-index)
          items-to-left (vec (take middle-index items))
          items-to-right (vec (drop (+ middle-index 1) items))]

      (cond
        (= item middle-item) (+ start-index middle-index)
        (< item middle-item) (recur items-to-left start-index)
        (> item middle-item) (recur items-to-right (+ 1 start-index middle-index))
        )
      )
    )
  )
