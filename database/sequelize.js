require('dotenv').config()

module.exports = {
	[process.env.NODE_ENV]: {
		dialect: process.env.DB_DIALECT,
		host: process.env.DB_HOST,
		port: +process.env.DB_PORT,
		username: process.env.DB_USERNAME,
		password: process.env.DB_PASSWORD,
		database: process.env.DB_NAME,
		dialectOptions: {
			ssl: JSON.parse(process.env.DB_SSL || 'false')
				? {
						require: JSON.parse(process.env.DB_SSL || 'false'),
						rejectUnauthorized: JSON.parse(process.env.DB_SSL || 'false')
				  }
				: JSON.parse(process.env.DB_SSL || 'false')
		}
	}
}
