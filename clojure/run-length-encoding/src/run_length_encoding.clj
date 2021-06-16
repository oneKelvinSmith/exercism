(ns run-length-encoding)


(defn encode [[character & others :as partition]]
  (cond->> character
           (some? others) (str (count partition))))

(defn run-length-encode
  "encodes a string with run-length-encoding"
  [plain-text]
  (->> plain-text
       (partition-by identity)
       (map encode)
       (apply str)))


(defn- decode [[_ number character]]
  (cond->> character
           (some? number) (repeat (Integer. number))
           true (apply str)))

(defn run-length-decode
  "decodes a run-length-encoded string"
  [cipher-text]
  (->> cipher-text
       (re-seq #"(\d+)?(\D)")
       (map decode)
       (apply str)))
