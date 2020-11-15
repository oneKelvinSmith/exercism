(ns bob,
  (:require [clojure.string :as string]))

(defn- silence? [remark] (string/blank? remark))
(defn- sensible? [remark] (not (nil? (re-find #"[A-z]" remark))))
(defn- question? [remark] (string/ends-with? remark "?"))
(defn- forceful? [remark] (and (sensible? remark) (= remark (string/upper-case remark))))
(defn- exclamation? [remark] (and (forceful? remark) (question? remark)))

(defn response-for [s]
  (let [remark (string/trim s)]
    (cond
      (silence? remark) "Fine. Be that way!"
      (exclamation? remark) "Calm down, I know what I'm doing!"
      (question? remark) "Sure."
      (forceful? remark) "Whoa, chill out!"
      :else "Whatever."
      ))
  )
