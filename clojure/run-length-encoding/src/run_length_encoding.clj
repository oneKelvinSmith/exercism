(ns run-length-encoding)

(defn run-length-encode
  "encodes a string with run-length-encoding"
  [plain-text]
  (let [character-groups (partition-by identity plain-text)
        group-encoder (fn [encoded group]
                        (let [character (first group)
                              frequency (count group)
                              code (if (= 1 frequency) "" frequency)]
                          (str encoded code character)))]
    (reduce group-encoder "" character-groups)))

(defn run-length-decode
  "decodes a run-length-encoded string"
  [cipher-text])
