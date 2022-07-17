/*
 * @Author: Esword
 * @Description:
 * @FileName:  js
 * @Version: 1.0.0
 * @Date: 2022-07-16 16:04
 */

package consts

const SecondaryDecryptScript = `
	function atob(r) {
		e = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/=";
		var o = String(r).replace(/=+$/, "");
		if (o.length % 4 == 1) throw new t("'atob' failed: The string to be decoded is not correctly encoded.");
		for (var n, a, i = 0, c = 0, d = ""; a = o.charAt(c++); ~a && (n = i % 4 ? 64 * n + a : a, i++ % 4) ? d += String.fromCharCode(255 & n >> (-2 * i & 6)) : 0) a = e.indexOf(a);
		return d
	}

	function btoa(r) {
		e = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/=";
		for (var o, n, a = String(r), i = 0, c = e, d = ""; a.charAt(0 | i) || (c = "=", i % 1); d += c.charAt(63 & o >> 8 - i % 1 * 8)) {
			if (n = a.charCodeAt(i += .75), n > 255) throw new t("'btoa' failed: The string to be encoded contains characters outside of the Latin1 range.");
			o = o << 8 | n
		}
		return d
	}
	
	function E(e) {
		var t = Array.prototype.map.call(e, function (e) {
			return String.fromCharCode(e)
		}).join("");
		return btoa(t)
	}
	
	function g(e) {
		var t = new ArrayBuffer(e.length);
		var n = new Uint8Array(t);
		for (var i = 0, a = e.length; i < a; i++)
			n[i] = e.charCodeAt(i);
		return n
	}
	
	function b(e) {
		var t = atob(e);
		var n = new Uint8Array(t.length);
		Array.prototype.forEach.call(t, function (e, t) {
			n[t] = e.charCodeAt(0)
		});
		return n
	}
	
	function c(e) {
		var t = new Uint8Array(e.match(/[\da-f]{2}/gi).map(function (e) {
			return parseInt(e, 16)
		}));
		var n = t.buffer;
		return n
	}
	
	function d(e) {
		var t = e.trim();
		var n = "0x" === t.substr(0, 2).toLowerCase() ? t.substr(2) : t;
		var i = n.length;
		if (i % 2 !== 0) {
			alert("Illegal Format ASCII Code!");
			return ""
		}
		var a;
		var o = [];
		for (var s = 0; s < i; s += 2) {
			a = parseInt(n.substr(s, 2), 16);
			o.push(String.fromCharCode(a))
		}
		return o.join("")
	}
	
	function p(e) {
		var t = "";
		var n = 8192;
		var i;
		for (i = 0; i < e.length / n; i++)
			t += String.fromCharCode.apply(null, e.slice(i * n, (i + 1) * n));
		t += String.fromCharCode.apply(null, e.slice(i * n));
		return t
	}
	
	function m(e) {
		var t = [], n, i, a;
		for (i = 0; i < e.length; ++i) {
			n = e[i];
			for (a = 3; a >= 0; --a)
				t.push(n >> 8 * a & 255)
		}
		return t
	}
`

const TokenSCRIPT = `	
		function atob(r) {
			e = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/=";
			var o = String(r).replace(/=+$/, "");
			if (o.length % 4 == 1) throw new t("'atob' failed: The string to be decoded is not correctly encoded.");
			for (var n, a, i = 0, c = 0, d = ""; a = o.charAt(c++); ~a && (n = i % 4 ? 64 * n + a : a, i++ % 4) ? d += String.fromCharCode(255 & n >> (-2 * i & 6)) : 0) a = e.indexOf(a);
			return d
		}
		
		function btoa(r) {
			e = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/=";
			for (var o, n, a = String(r), i = 0, c = e, d = ""; a.charAt(0 | i) || (c = "=", i % 1); d += c.charAt(63 & o >> 8 - i % 1 * 8)) {
				if (n = a.charCodeAt(i += .75), n > 255) throw new t("'btoa' failed: The string to be encoded contains characters outside of the Latin1 range.");
				o = o << 8 | n
			}
			return d
		}
		
		i = {
			"charListMap": {
				"v_0": [
					{
						"t": 1,
						"i": 2
					},
					{
						"t": 3,
						"i": 9
					},
					{
						"t": 3,
						"i": 7
					},
					{
						"t": 3,
						"i": 2
					},
					{
						"t": 1,
						"i": 0
					},
					{
						"t": 1,
						"i": 4
					},
					{
						"t": 1,
						"i": 2
					},
					{
						"t": 3,
						"i": 1
					},
					{
						"t": 3,
						"i": 1
					},
					{
						"t": 3,
						"i": 8
					},
					{
						"t": 1,
						"i": 5
					},
					{
						"t": 1,
						"i": 4
					},
					{
						"t": 1,
						"i": 4
					},
					{
						"t": 3,
						"i": 6
					},
					{
						"t": 3,
						"i": 9
					},
					{
						"t": 1,
						"i": 5
					}
				],
				"v_1": [
					{
						"t": 3,
						"i": 3
					},
					{
						"t": 1,
						"i": 5
					},
					{
						"t": 1,
						"i": 15
					},
					{
						"t": 3,
						"i": 4
					},
					{
						"t": 1,
						"i": 23
					},
					{
						"t": 1,
						"i": 18
					},
					{
						"t": 3,
						"i": 9
					},
					{
						"t": 3,
						"i": 2
					},
					{
						"t": 3,
						"i": 2
					},
					{
						"t": 1,
						"i": 14
					},
					{
						"t": 1,
						"i": 20
					},
					{
						"t": 1,
						"i": 22
					},
					{
						"t": 3,
						"i": 5
					},
					{
						"t": 1,
						"i": 16
					},
					{
						"t": 3,
						"i": 7
					},
					{
						"t": 3,
						"i": 2
					}
				],
				"v_2": [
					{
						"t": 2,
						"i": 16
					},
					{
						"t": 2,
						"i": 7
					},
					{
						"t": 1,
						"i": 7
					},
					{
						"t": 2,
						"i": 24
					},
					{
						"t": 1,
						"i": 17
					},
					{
						"t": 2,
						"i": 4
					},
					{
						"t": 1,
						"i": 4
					},
					{
						"t": 2,
						"i": 18
					},
					{
						"t": 2,
						"i": 12
					},
					{
						"t": 2,
						"i": 5
					},
					{
						"t": 2,
						"i": 18
					},
					{
						"t": 2,
						"i": 4
					},
					{
						"t": 1,
						"i": 0
					},
					{
						"t": 2,
						"i": 22
					},
					{
						"t": 1,
						"i": 11
					},
					{
						"t": 2,
						"i": 6
					}
				],
				"v_3": [
					{
						"t": 2,
						"i": 18
					},
					{
						"t": 1,
						"i": 4
					},
					{
						"t": 1,
						"i": 7
					},
					{
						"t": 2,
						"i": 24
					},
					{
						"t": 1,
						"i": 17
					},
					{
						"t": 2,
						"i": 15
					},
					{
						"t": 1,
						"i": 4
					},
					{
						"t": 2,
						"i": 18
					},
					{
						"t": 1,
						"i": 11
					},
					{
						"t": 2,
						"i": 5
					},
					{
						"t": 2,
						"i": 18
					},
					{
						"t": 1,
						"i": 14
					},
					{
						"t": 1,
						"i": 0
					},
					{
						"t": 2,
						"i": 22
					},
					{
						"t": 1,
						"i": 11
					},
					{
						"t": 3,
						"i": 5
					}
				]
			},
			"charEnlist1": [
				"a",
				"b",
				"c",
				"d",
				"e",
				"f",
				"g",
				"h",
				"i",
				"j",
				"k",
				"l",
				"m",
				"n",
				"o",
				"p",
				"q",
				"r",
				"s",
				"t",
				"u",
				"v",
				"w",
				"x",
				"y",
				"z"
			],
			"charEnlist2": [
				"A",
				"B",
				"C",
				"D",
				"E",
				"F",
				"G",
				"H",
				"I",
				"J",
				"K",
				"L",
				"M",
				"N",
				"O",
				"P",
				"Q",
				"R",
				"S",
				"T",
				"U",
				"V",
				"W",
				"X",
				"Y",
				"Z"
			],
			"charEnlist3": [
				"0",
				"1",
				"2",
				"3",
				"4",
				"5",
				"6",
				"7",
				"8",
				"9"
			]
		}
		
		function _(e, t) {
			return l(t)[e]
		}
		
		function l(e) {
			return i["charEnlist" + e]
		}
		
		function u(e) {
			var t = i.charListMap["v_" + e]
				, n = "";
			for (var a = 0; a < t.length; a++)
				n += _(t[a].i, t[a].t);
			return n
		}
		
		function b(e) {
			var t = Array.prototype.map.call(e, function (e) {
				return String.fromCharCode(e)
			}).join("");
			return btoa(t)
		}
		
		function p(e) {
			var t = "";
			var n = 8192;
			var i;
			for (i = 0; i < e.length / n; i++)
				t += String.fromCharCode.apply(null, e.slice(i * n, (i + 1) * n));
			t += String.fromCharCode.apply(null, e.slice(i * n));
			return t
		}
		
		function m(e) {
			var t = [], n, i, a;
			for (i = 0; i < e.length; ++i) {
				n = e[i];
				for (a = 3; a >= 0; --a)
					t.push(n >> 8 * a & 255)
			}
			return t
		}
		
		function E(e) {
			var t = new ArrayBuffer(e.length);
			var n = new Uint8Array(t);
			for (var i = 0, a = e.length; i < a; i++)
				n[i] = e.charCodeAt(i);
			return n
		}
		
		function g(e) {
			var t = atob(e);
			var n = new Uint8Array(t.length);
			Array.prototype.forEach.call(t, function (e, t) {
				n[t] = e.charCodeAt(0)
			});
			return n
		}
		`
