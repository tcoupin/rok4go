const HtmlWebpackPlugin = require('html-webpack-plugin');
const path = require('path');

module.exports = {
	entry: './src/js/main.js',
	output: {
		filename: 'js/main.js',
		path: path.resolve(__dirname, 'dist')
	},
	module: {
		rules: [{
			test: /\.(css)$/,
			use: ['style-loader', 'css-loader']
		}, {
			test: /\.html$/,
			use: [ 
				{
					loader: "file-loader",
					options: {
						name: "[name].[ext]"
					}
				},{
					loader: "extract-loader"
				},{
					loader: "html-loader"
				}
			]
		}]
	}
};