(ns run-length-encoding)

(defn- encode [[character & _ :as partition]]
  (if (= 1 (count partition))
    character
    (str (count partition) character)))

(defn run-length-encode
  "encodes a string with run-length-encoding"
  [plain-text]
  (->> plain-text
       (partition-by identity)
       (map encode)
       (apply str)))

(defn- decode [[_ number character]]
  (->> character
       (repeat ((fnil #(Integer. %) 1) number))
       (apply str)))

(defn run-length-decode
  "decodes a run-length-encoded string"
  [cipher-text]
  (->> cipher-text
       (re-seq #"(\d+)?(\D)")
       (map decode)
       (apply str)))
