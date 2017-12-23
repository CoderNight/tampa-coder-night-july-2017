(ns missing.core
  (:use [clojure.string :only (split join)]
        clojure.java.io)
  (:gen-class))

(defn open-file
  "Opens the given file and returns an seq of the lines"
  [path]
  (with-open [r (reader path)]
    (doall (line-seq r))))

(defn printMissing
  "Takes a map of frequencies and prints positive frequencies"
  [keyvals]
  (join " " (map first keyvals)))

(defn detectMissing
  "given known frequencies, calculate remaining/missing from found numbers"
  [freq found]
  (reduce (fn [freqs item]
    (update-in freqs [item] dec)) freq found))

(defn listFoundNumbers
  "given a string, split into numbers"
  [strlist]
  (map #(Integer/parseInt %) (split strlist #" ")))

(defn findFrequencies
  "given a list of expected numbers, count frequencies of numbers"
  [strlist]
  (let [numbers (listFoundNumbers strlist)]
    (frequencies numbers)))

(defn findMissing
  "given a file location, find missing numbers"
  [filename]
  (let [contents (open-file filename)]
    (printMissing
      (filter (fn [[k v]] (> v 0))
      (detectMissing
        (findFrequencies (nth contents 3))
        (listFoundNumbers (nth contents 1)))))))

(defn -main
  "opens a file, and uses it to find missing numbers"
  [path & args]
  (println (findMissing path)))