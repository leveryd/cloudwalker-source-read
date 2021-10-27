package WebshellDetector

import "github.com/CyrusF/go-bayesian"

/*
WebshellDetector - Refactor version 1
Date	0814
Author	Twice
Intro	predict webshell by words using Naive Bayes
*/

type words struct {
	data []string
}

/*
算法模型：朴素贝叶斯
特征：代码ast的字符串。比如 <?php $a=1;var_dump($a+1);?>，会提取出 ["a","a","var_dump"]
*/
func (words words) Predict(model *bayesian.Classifier) float64 {
	score, _, _ := model.Classify(words.data...)
	return score["webshell"] / (score["webshell"] + score["normal"])
}
