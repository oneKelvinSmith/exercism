(ns binary-search)

(defn spy [tag value] (prn tag value) value)

(defn middle [collection]
  (quot (count collection) 2))

(defn search-for
  [item collection]
  (loop [items (spy "collection" (into [] collection))
         start-index 0]
    (if (empty? items) (throw (Exception. "not found")))

    (let [middle-index (middle items)
          middle-item (get items middle-index)
          items-to-left (into [] (take middle-index items))
          items-to-right (into [] (drop (+ middle-index 1) items))]

      (cond
        (= item middle-item) (spy "equal" (+ start-index middle-index))
        (< item middle-item) (recur (spy "left" items-to-left) start-index)
        (> item middle-item) (recur (spy "right" items-to-right) (+ 1 start-index middle-index))
        )
      )
    )
  )
