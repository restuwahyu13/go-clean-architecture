import { QueryInterface, Sequelize } from 'sequelize'
import bcrypt from 'bcrypt'

module.exports = {
	up: async (queryInterface: QueryInterface, _sequelize: Sequelize) => {
		const hashPassword: string = bcrypt.hashSync('@Qwerty12', 12)
		const users: Record<string, any>[] = [
			{
				name: 'Mat Metal',
				email: 'matmetal13@gmail.com',
				password: hashPassword
			},
			{
				name: 'Anto Killer',
				email: 'antokiller13@gmail.com',
				password: hashPassword
			},
			{
				name: 'Jamal Cavalera',
				email: 'jamal13@gmail.com',
				password: hashPassword
			},
			{
				name: 'Santoso',
				email: 'santoso13@gmail.com',
				password: hashPassword
			}
		]

		return queryInterface.bulkInsert('users', users, { logging: true })
	},
	down: async (queryInterface: QueryInterface, _sequelize: Sequelize) => {
		return queryInterface.bulkDelete('users', { logging: true })
	}
}
