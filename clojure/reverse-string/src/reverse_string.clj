(ns reverse-string
  (:require [clojure.string :as str]))

(defn reverse-string [s]
  (str/join (reverse s)))
