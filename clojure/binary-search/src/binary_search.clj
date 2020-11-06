(ns binary-search)

(defn spy [tag value] (prn tag value) value)

(defn middle [collection]
  (quot (count collection) 2))

(defn search-for
  ([item collection]
   (loop [items collection
          index 0]
     (let [middle-index (middle items)
           middle-item (get items middle-index)
           items-to-left (into [] (take middle-index items))
           items-to-right (into [] (drop (+ middle-index 1) items))]

       (cond
         (< item middle-item) (recur items-to-left index)
         (> item middle-item) (recur items-to-right (+ 1 index middle-index))
         :else (+ index middle-index))
       )))
  )
