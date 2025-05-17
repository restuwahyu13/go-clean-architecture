import { QueryInterface, Sequelize, DataTypes } from 'sequelize'

module.exports = {
	up: async (queryInterface: QueryInterface, sequelize: Sequelize) => {
		const tablExist: boolean = await queryInterface.tableExists('users')
		if (!tablExist) {
			await queryInterface.createTable(
				'users',
				{
					id: { type: DataTypes.UUID, primaryKey: true, allowNull: false, unique: true, defaultValue: sequelize.literal('uuid_generate_v4()') },
					name: { type: DataTypes.STRING(200), allowNull: false },
					email: { type: DataTypes.STRING(50), allowNull: false },
					status: { type: DataTypes.STRING(25), allowNull: false, defaultValue: 'active' },
					password: { type: DataTypes.TEXT, allowNull: false },
					created_at: { type: DataTypes.DATE, allowNull: false, defaultValue: sequelize.literal('CURRENT_TIMESTAMP') },
					updated_at: { type: DataTypes.DATE },
					deleted_at: { type: DataTypes.DATE }
				},
				{
					logging: true
				}
			)
		}
	},
	down: async (queryInterface: QueryInterface, _sequelize: Sequelize) => {
		const tableExist: boolean = await queryInterface.tableExists('users')
		if (tableExist) {
			return queryInterface.dropTable('users')
		}
	}
}
