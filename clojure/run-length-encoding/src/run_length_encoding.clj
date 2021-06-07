(ns run-length-encoding
  (:require [clojure.string :as string]))

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
       (string/join)))

(defn- decode [[_ number character]]
  (if (nil? number)
    character
    (string/join (repeat (Integer. number) character))))

(defn run-length-decode
  "decodes a run-length-encoded string"
  [cipher-text]
  (->> cipher-text
       (re-seq #"(\d+)?(\D)")
       (map decode)
       (string/join)))
