package second_solution

import (
	"reflect"
	"testing"
)

func TestCodec_serializeAndDeserialize(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
	}{
		{
			name: "Simple tree",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   3,
					Left:  nil,
					Right: nil,
				},
			},
		},
		{
			name: "Tree without right node",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
		},
		{
			name: "Tree without left node",
			root: &TreeNode{
				Val:  1,
				Left: nil,
				Right: &TreeNode{
					Val:   2,
					Left:  nil,
					Right: nil,
				},
			},
		},
		{
			name: "nil",
			root: nil,
		},
		{
			name: "only root",
			root: &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
		},
		{
			name: "Big tree",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 1,
						Left: &TreeNode{
							Val:  2,
							Left: nil,
							Right: &TreeNode{
								Val: 1000,
								Left: &TreeNode{
									Val:   444,
									Left:  nil,
									Right: nil,
								},
								Right: &TreeNode{
									Val:  7777777,
									Left: nil,
									Right: &TreeNode{
										Val: 88888888,
										Left: &TreeNode{
											Val: 999999999,
											Left: &TreeNode{
												Val:  123456789,
												Left: nil,
												Right: &TreeNode{
													Val: -777,
													Left: &TreeNode{
														Val:   -666,
														Left:  nil,
														Right: nil,
													},
													Right: &TreeNode{
														Val: -1,
														Left: &TreeNode{
															Val:   0,
															Left:  nil,
															Right: nil,
														},
														Right: nil,
													},
												},
											},
											Right: nil,
										},
										Right: nil,
									},
								},
							},
						},
						Right: &TreeNode{
							Val:  123,
							Left: nil,
							Right: &TreeNode{
								Val:  456,
								Left: nil,
								Right: &TreeNode{
									Val: 1,
									Left: &TreeNode{
										Val:  2,
										Left: nil,
										Right: &TreeNode{
											Val: 333,
											Left: &TreeNode{
												Val:   444,
												Left:  nil,
												Right: nil,
											},
											Right: &TreeNode{
												Val:  7777777,
												Left: nil,
												Right: &TreeNode{
													Val: 88888888,
													Left: &TreeNode{
														Val: 999999999,
														Left: &TreeNode{
															Val:  123456789,
															Left: nil,
															Right: &TreeNode{
																Val: -777,
																Left: &TreeNode{
																	Val:   -666,
																	Left:  nil,
																	Right: nil,
																},
																Right: &TreeNode{
																	Val: -1,
																	Left: &TreeNode{
																		Val:   0,
																		Left:  nil,
																		Right: nil,
																	},
																	Right: nil,
																},
															},
														},
														Right: nil,
													},
													Right: nil,
												},
											},
										},
									},
									Right: &TreeNode{
										Val:  123,
										Left: nil,
										Right: &TreeNode{
											Val:  456,
											Left: nil,
											Right: &TreeNode{
												Val: 1,
												Left: &TreeNode{
													Val:  2,
													Left: nil,
													Right: &TreeNode{
														Val: 333,
														Left: &TreeNode{
															Val:   444,
															Left:  nil,
															Right: nil,
														},
														Right: &TreeNode{
															Val:  7777777,
															Left: nil,
															Right: &TreeNode{
																Val: 88888888,
																Left: &TreeNode{
																	Val: 999999999,
																	Left: &TreeNode{
																		Val:  123456789,
																		Left: nil,
																		Right: &TreeNode{
																			Val: -777,
																			Left: &TreeNode{
																				Val:   -666,
																				Left:  nil,
																				Right: nil,
																			},
																			Right: &TreeNode{
																				Val: -1,
																				Left: &TreeNode{
																					Val:   0,
																					Left:  nil,
																					Right: nil,
																				},
																				Right: nil,
																			},
																		},
																	},
																	Right: nil,
																},
																Right: nil,
															},
														},
													},
												},
												Right: &TreeNode{
													Val:  123,
													Left: nil,
													Right: &TreeNode{
														Val:  456,
														Left: nil,
														Right: &TreeNode{
															Val: 1,
															Left: &TreeNode{
																Val:  2,
																Left: nil,
																Right: &TreeNode{
																	Val: 333,
																	Left: &TreeNode{
																		Val:   444,
																		Left:  nil,
																		Right: nil,
																	},
																	Right: &TreeNode{
																		Val:  7777777,
																		Left: nil,
																		Right: &TreeNode{
																			Val: 88888888,
																			Left: &TreeNode{
																				Val: 999999999,
																				Left: &TreeNode{
																					Val:  123456789,
																					Left: nil,
																					Right: &TreeNode{
																						Val: -777,
																						Left: &TreeNode{
																							Val:   -666,
																							Left:  nil,
																							Right: nil,
																						},
																						Right: &TreeNode{
																							Val: -1,
																							Left: &TreeNode{
																								Val:   0,
																								Left:  nil,
																								Right: nil,
																							},
																							Right: nil,
																						},
																					},
																				},
																				Right: nil,
																			},
																			Right: nil,
																		},
																	},
																},
															},
															Right: &TreeNode{
																Val:  123,
																Left: nil,
																Right: &TreeNode{
																	Val:  456,
																	Left: nil,
																	Right: &TreeNode{
																		Val: 1,
																		Left: &TreeNode{
																			Val:  2,
																			Left: nil,
																			Right: &TreeNode{
																				Val: 333,
																				Left: &TreeNode{
																					Val:   444,
																					Left:  nil,
																					Right: nil,
																				},
																				Right: &TreeNode{
																					Val:  7777777,
																					Left: nil,
																					Right: &TreeNode{
																						Val: 88888888,
																						Left: &TreeNode{
																							Val: 999999999,
																							Left: &TreeNode{
																								Val:  123456789,
																								Left: nil,
																								Right: &TreeNode{
																									Val: -777,
																									Left: &TreeNode{
																										Val:   -666,
																										Left:  nil,
																										Right: nil,
																									},
																									Right: &TreeNode{
																										Val: -1,
																										Left: &TreeNode{
																											Val:   0,
																											Left:  nil,
																											Right: nil,
																										},
																										Right: nil,
																									},
																								},
																							},
																							Right: nil,
																						},
																						Right: nil,
																					},
																				},
																			},
																		},
																		Right: &TreeNode{
																			Val:  123,
																			Left: nil,
																			Right: &TreeNode{
																				Val: 456,
																				Left: &TreeNode{
																					Val: 1,
																					Left: &TreeNode{
																						Val:  2,
																						Left: nil,
																						Right: &TreeNode{
																							Val: 1000,
																							Left: &TreeNode{
																								Val:   444,
																								Left:  nil,
																								Right: nil,
																							},
																							Right: &TreeNode{
																								Val:  7777777,
																								Left: nil,
																								Right: &TreeNode{
																									Val: 88888888,
																									Left: &TreeNode{
																										Val: 999999999,
																										Left: &TreeNode{
																											Val:  123456789,
																											Left: nil,
																											Right: &TreeNode{
																												Val: -777,
																												Left: &TreeNode{
																													Val:   -666,
																													Left:  nil,
																													Right: nil,
																												},
																												Right: &TreeNode{
																													Val: -1,
																													Left: &TreeNode{
																														Val:   0,
																														Left:  nil,
																														Right: nil,
																													},
																													Right: nil,
																												},
																											},
																										},
																										Right: nil,
																									},
																									Right: nil,
																								},
																							},
																						},
																					},
																					Right: &TreeNode{
																						Val:  123,
																						Left: nil,
																						Right: &TreeNode{
																							Val:  456,
																							Left: nil,
																							Right: &TreeNode{
																								Val: 1,
																								Left: &TreeNode{
																									Val:  2,
																									Left: nil,
																									Right: &TreeNode{
																										Val: 333,
																										Left: &TreeNode{
																											Val:   444,
																											Left:  nil,
																											Right: nil,
																										},
																										Right: &TreeNode{
																											Val:  7777777,
																											Left: nil,
																											Right: &TreeNode{
																												Val: 88888888,
																												Left: &TreeNode{
																													Val: 999999999,
																													Left: &TreeNode{
																														Val:  123456789,
																														Left: nil,
																														Right: &TreeNode{
																															Val: -777,
																															Left: &TreeNode{
																																Val:   -666,
																																Left:  nil,
																																Right: nil,
																															},
																															Right: &TreeNode{
																																Val: -1,
																																Left: &TreeNode{
																																	Val:   0,
																																	Left:  nil,
																																	Right: nil,
																																},
																																Right: nil,
																															},
																														},
																													},
																													Right: nil,
																												},
																												Right: nil,
																											},
																										},
																									},
																								},
																								Right: &TreeNode{
																									Val:  123,
																									Left: nil,
																									Right: &TreeNode{
																										Val:  456,
																										Left: nil,
																										Right: &TreeNode{
																											Val: 1,
																											Left: &TreeNode{
																												Val:  2,
																												Left: nil,
																												Right: &TreeNode{
																													Val: 333,
																													Left: &TreeNode{
																														Val:   444,
																														Left:  nil,
																														Right: nil,
																													},
																													Right: &TreeNode{
																														Val:  7777777,
																														Left: nil,
																														Right: &TreeNode{
																															Val: 88888888,
																															Left: &TreeNode{
																																Val: 999999999,
																																Left: &TreeNode{
																																	Val:  123456789,
																																	Left: nil,
																																	Right: &TreeNode{
																																		Val: -777,
																																		Left: &TreeNode{
																																			Val:   -666,
																																			Left:  nil,
																																			Right: nil,
																																		},
																																		Right: &TreeNode{
																																			Val: -1,
																																			Left: &TreeNode{
																																				Val:   0,
																																				Left:  nil,
																																				Right: nil,
																																			},
																																			Right: nil,
																																		},
																																	},
																																},
																																Right: nil,
																															},
																															Right: nil,
																														},
																													},
																												},
																											},
																											Right: &TreeNode{
																												Val:  123,
																												Left: nil,
																												Right: &TreeNode{
																													Val:  456,
																													Left: nil,
																													Right: &TreeNode{
																														Val: 1,
																														Left: &TreeNode{
																															Val:  2,
																															Left: nil,
																															Right: &TreeNode{
																																Val: 333,
																																Left: &TreeNode{
																																	Val:   444,
																																	Left:  nil,
																																	Right: nil,
																																},
																																Right: &TreeNode{
																																	Val:  7777777,
																																	Left: nil,
																																	Right: &TreeNode{
																																		Val: 88888888,
																																		Left: &TreeNode{
																																			Val: 999999999,
																																			Left: &TreeNode{
																																				Val:  123456789,
																																				Left: nil,
																																				Right: &TreeNode{
																																					Val: -777,
																																					Left: &TreeNode{
																																						Val:   -666,
																																						Left:  nil,
																																						Right: nil,
																																					},
																																					Right: &TreeNode{
																																						Val: -1,
																																						Left: &TreeNode{
																																							Val:   0,
																																							Left:  nil,
																																							Right: nil,
																																						},
																																						Right: nil,
																																					},
																																				},
																																			},
																																			Right: nil,
																																		},
																																		Right: nil,
																																	},
																																},
																															},
																														},
																														Right: &TreeNode{
																															Val:  123,
																															Left: nil,
																															Right: &TreeNode{
																																Val:  456,
																																Left: nil,
																																Right: &TreeNode{
																																	Val: 1,
																																	Left: &TreeNode{
																																		Val:  2,
																																		Left: nil,
																																		Right: &TreeNode{
																																			Val: 333,
																																			Left: &TreeNode{
																																				Val:   444,
																																				Left:  nil,
																																				Right: nil,
																																			},
																																			Right: &TreeNode{
																																				Val:  7777777,
																																				Left: nil,
																																				Right: &TreeNode{
																																					Val: 88888888,
																																					Left: &TreeNode{
																																						Val: 999999999,
																																						Left: &TreeNode{
																																							Val:  123456789,
																																							Left: nil,
																																							Right: &TreeNode{
																																								Val: -777,
																																								Left: &TreeNode{
																																									Val:   -666,
																																									Left:  nil,
																																									Right: nil,
																																								},
																																								Right: &TreeNode{
																																									Val: -1,
																																									Left: &TreeNode{
																																										Val:   0,
																																										Left:  nil,
																																										Right: nil,
																																									},
																																									Right: nil,
																																								},
																																							},
																																						},
																																						Right: nil,
																																					},
																																					Right: nil,
																																				},
																																			},
																																		},
																																	},
																																	Right: &TreeNode{
																																		Val:  123,
																																		Left: nil,
																																		Right: &TreeNode{
																																			Val:   456,
																																			Left:  nil,
																																			Right: nil,
																																		},
																																	},
																																},
																															},
																														},
																													},
																												},
																											},
																										},
																									},
																								},
																							},
																						},
																					},
																				},
																				Right: &TreeNode{
																					Val: 1,
																					Left: &TreeNode{
																						Val:  2,
																						Left: nil,
																						Right: &TreeNode{
																							Val: 1000,
																							Left: &TreeNode{
																								Val:   444,
																								Left:  nil,
																								Right: nil,
																							},
																							Right: &TreeNode{
																								Val:  7777777,
																								Left: nil,
																								Right: &TreeNode{
																									Val: 88888888,
																									Left: &TreeNode{
																										Val: 999999999,
																										Left: &TreeNode{
																											Val:  123456789,
																											Left: nil,
																											Right: &TreeNode{
																												Val: -777,
																												Left: &TreeNode{
																													Val:   -666,
																													Left:  nil,
																													Right: nil,
																												},
																												Right: &TreeNode{
																													Val: -1,
																													Left: &TreeNode{
																														Val:   0,
																														Left:  nil,
																														Right: nil,
																													},
																													Right: nil,
																												},
																											},
																										},
																										Right: nil,
																									},
																									Right: nil,
																								},
																							},
																						},
																					},
																					Right: &TreeNode{
																						Val:  123,
																						Left: nil,
																						Right: &TreeNode{
																							Val:  456,
																							Left: nil,
																							Right: &TreeNode{
																								Val: 1,
																								Left: &TreeNode{
																									Val:  2,
																									Left: nil,
																									Right: &TreeNode{
																										Val: 333,
																										Left: &TreeNode{
																											Val:   444,
																											Left:  nil,
																											Right: nil,
																										},
																										Right: &TreeNode{
																											Val:  7777777,
																											Left: nil,
																											Right: &TreeNode{
																												Val: 88888888,
																												Left: &TreeNode{
																													Val: 999999999,
																													Left: &TreeNode{
																														Val:  123456789,
																														Left: nil,
																														Right: &TreeNode{
																															Val: -777,
																															Left: &TreeNode{
																																Val:   -666,
																																Left:  nil,
																																Right: nil,
																															},
																															Right: &TreeNode{
																																Val: -1,
																																Left: &TreeNode{
																																	Val:   0,
																																	Left:  nil,
																																	Right: nil,
																																},
																																Right: nil,
																															},
																														},
																													},
																													Right: nil,
																												},
																												Right: nil,
																											},
																										},
																									},
																								},
																								Right: &TreeNode{
																									Val:  123,
																									Left: nil,
																									Right: &TreeNode{
																										Val:  456,
																										Left: nil,
																										Right: &TreeNode{
																											Val: 1,
																											Left: &TreeNode{
																												Val:  2,
																												Left: nil,
																												Right: &TreeNode{
																													Val: 333,
																													Left: &TreeNode{
																														Val:   444,
																														Left:  nil,
																														Right: nil,
																													},
																													Right: &TreeNode{
																														Val:  7777777,
																														Left: nil,
																														Right: &TreeNode{
																															Val: 88888888,
																															Left: &TreeNode{
																																Val: 999999999,
																																Left: &TreeNode{
																																	Val:  123456789,
																																	Left: nil,
																																	Right: &TreeNode{
																																		Val: -777,
																																		Left: &TreeNode{
																																			Val:   -666,
																																			Left:  nil,
																																			Right: nil,
																																		},
																																		Right: &TreeNode{
																																			Val: -1,
																																			Left: &TreeNode{
																																				Val:   0,
																																				Left:  nil,
																																				Right: nil,
																																			},
																																			Right: nil,
																																		},
																																	},
																																},
																																Right: nil,
																															},
																															Right: nil,
																														},
																													},
																												},
																											},
																											Right: &TreeNode{
																												Val:  123,
																												Left: nil,
																												Right: &TreeNode{
																													Val:  456,
																													Left: nil,
																													Right: &TreeNode{
																														Val: 1,
																														Left: &TreeNode{
																															Val:  2,
																															Left: nil,
																															Right: &TreeNode{
																																Val: 333,
																																Left: &TreeNode{
																																	Val:   444,
																																	Left:  nil,
																																	Right: nil,
																																},
																																Right: &TreeNode{
																																	Val:  7777777,
																																	Left: nil,
																																	Right: &TreeNode{
																																		Val: 88888888,
																																		Left: &TreeNode{
																																			Val: 999999999,
																																			Left: &TreeNode{
																																				Val:  123456789,
																																				Left: nil,
																																				Right: &TreeNode{
																																					Val: -777,
																																					Left: &TreeNode{
																																						Val:   -666,
																																						Left:  nil,
																																						Right: nil,
																																					},
																																					Right: &TreeNode{
																																						Val: -1,
																																						Left: &TreeNode{
																																							Val:   0,
																																							Left:  nil,
																																							Right: nil,
																																						},
																																						Right: nil,
																																					},
																																				},
																																			},
																																			Right: nil,
																																		},
																																		Right: nil,
																																	},
																																},
																															},
																														},
																														Right: &TreeNode{
																															Val:  123,
																															Left: nil,
																															Right: &TreeNode{
																																Val:  456,
																																Left: nil,
																																Right: &TreeNode{
																																	Val: 1,
																																	Left: &TreeNode{
																																		Val:  2,
																																		Left: nil,
																																		Right: &TreeNode{
																																			Val: 333,
																																			Left: &TreeNode{
																																				Val:   444,
																																				Left:  nil,
																																				Right: nil,
																																			},
																																			Right: &TreeNode{
																																				Val:  7777777,
																																				Left: nil,
																																				Right: &TreeNode{
																																					Val: 88888888,
																																					Left: &TreeNode{
																																						Val: 999999999,
																																						Left: &TreeNode{
																																							Val:  123456789,
																																							Left: nil,
																																							Right: &TreeNode{
																																								Val: -777,
																																								Left: &TreeNode{
																																									Val:   -666,
																																									Left:  nil,
																																									Right: nil,
																																								},
																																								Right: &TreeNode{
																																									Val: -1,
																																									Left: &TreeNode{
																																										Val:   0,
																																										Left:  nil,
																																										Right: nil,
																																									},
																																									Right: nil,
																																								},
																																							},
																																						},
																																						Right: nil,
																																					},
																																					Right: nil,
																																				},
																																			},
																																		},
																																	},
																																	Right: &TreeNode{
																																		Val:  123,
																																		Left: nil,
																																		Right: &TreeNode{
																																			Val: 456,
																																			Left: &TreeNode{
																																				Val: 1,
																																				Left: &TreeNode{
																																					Val:  2,
																																					Left: nil,
																																					Right: &TreeNode{
																																						Val: 1000,
																																						Left: &TreeNode{
																																							Val:   444,
																																							Left:  nil,
																																							Right: nil,
																																						},
																																						Right: &TreeNode{
																																							Val:  7777777,
																																							Left: nil,
																																							Right: &TreeNode{
																																								Val: 88888888,
																																								Left: &TreeNode{
																																									Val: 999999999,
																																									Left: &TreeNode{
																																										Val:  123456789,
																																										Left: nil,
																																										Right: &TreeNode{
																																											Val: -777,
																																											Left: &TreeNode{
																																												Val:   -666,
																																												Left:  nil,
																																												Right: nil,
																																											},
																																											Right: &TreeNode{
																																												Val: -1,
																																												Left: &TreeNode{
																																													Val:   0,
																																													Left:  nil,
																																													Right: nil,
																																												},
																																												Right: nil,
																																											},
																																										},
																																									},
																																									Right: nil,
																																								},
																																								Right: nil,
																																							},
																																						},
																																					},
																																				},
																																				Right: &TreeNode{
																																					Val:  123,
																																					Left: nil,
																																					Right: &TreeNode{
																																						Val:  456,
																																						Left: nil,
																																						Right: &TreeNode{
																																							Val: 1,
																																							Left: &TreeNode{
																																								Val:  2,
																																								Left: nil,
																																								Right: &TreeNode{
																																									Val: 333,
																																									Left: &TreeNode{
																																										Val:   444,
																																										Left:  nil,
																																										Right: nil,
																																									},
																																									Right: &TreeNode{
																																										Val:  7777777,
																																										Left: nil,
																																										Right: &TreeNode{
																																											Val: 88888888,
																																											Left: &TreeNode{
																																												Val: 999999999,
																																												Left: &TreeNode{
																																													Val:  123456789,
																																													Left: nil,
																																													Right: &TreeNode{
																																														Val: -777,
																																														Left: &TreeNode{
																																															Val:   -666,
																																															Left:  nil,
																																															Right: nil,
																																														},
																																														Right: &TreeNode{
																																															Val: -1,
																																															Left: &TreeNode{
																																																Val:   0,
																																																Left:  nil,
																																																Right: nil,
																																															},
																																															Right: nil,
																																														},
																																													},
																																												},
																																												Right: nil,
																																											},
																																											Right: nil,
																																										},
																																									},
																																								},
																																							},
																																							Right: &TreeNode{
																																								Val:  123,
																																								Left: nil,
																																								Right: &TreeNode{
																																									Val:  456,
																																									Left: nil,
																																									Right: &TreeNode{
																																										Val: 1,
																																										Left: &TreeNode{
																																											Val:  2,
																																											Left: nil,
																																											Right: &TreeNode{
																																												Val: 333,
																																												Left: &TreeNode{
																																													Val:   444,
																																													Left:  nil,
																																													Right: nil,
																																												},
																																												Right: &TreeNode{
																																													Val:  7777777,
																																													Left: nil,
																																													Right: &TreeNode{
																																														Val: 88888888,
																																														Left: &TreeNode{
																																															Val: 999999999,
																																															Left: &TreeNode{
																																																Val:  123456789,
																																																Left: nil,
																																																Right: &TreeNode{
																																																	Val: -777,
																																																	Left: &TreeNode{
																																																		Val:   -666,
																																																		Left:  nil,
																																																		Right: nil,
																																																	},
																																																	Right: &TreeNode{
																																																		Val: -1,
																																																		Left: &TreeNode{
																																																			Val:   0,
																																																			Left:  nil,
																																																			Right: nil,
																																																		},
																																																		Right: nil,
																																																	},
																																																},
																																															},
																																															Right: nil,
																																														},
																																														Right: nil,
																																													},
																																												},
																																											},
																																										},
																																										Right: &TreeNode{
																																											Val:  123,
																																											Left: nil,
																																											Right: &TreeNode{
																																												Val:  456,
																																												Left: nil,
																																												Right: &TreeNode{
																																													Val: 1,
																																													Left: &TreeNode{
																																														Val:  2,
																																														Left: nil,
																																														Right: &TreeNode{
																																															Val: 333,
																																															Left: &TreeNode{
																																																Val:   444,
																																																Left:  nil,
																																																Right: nil,
																																															},
																																															Right: &TreeNode{
																																																Val:  7777777,
																																																Left: nil,
																																																Right: &TreeNode{
																																																	Val: 88888888,
																																																	Left: &TreeNode{
																																																		Val: 999999999,
																																																		Left: &TreeNode{
																																																			Val:  123456789,
																																																			Left: nil,
																																																			Right: &TreeNode{
																																																				Val: -777,
																																																				Left: &TreeNode{
																																																					Val:   -666,
																																																					Left:  nil,
																																																					Right: nil,
																																																				},
																																																				Right: &TreeNode{
																																																					Val: -1,
																																																					Left: &TreeNode{
																																																						Val:   0,
																																																						Left:  nil,
																																																						Right: nil,
																																																					},
																																																					Right: nil,
																																																				},
																																																			},
																																																		},
																																																		Right: nil,
																																																	},
																																																	Right: nil,
																																																},
																																															},
																																														},
																																													},
																																													Right: &TreeNode{
																																														Val:  123,
																																														Left: nil,
																																														Right: &TreeNode{
																																															Val:  456,
																																															Left: nil,
																																															Right: &TreeNode{
																																																Val: 1,
																																																Left: &TreeNode{
																																																	Val:  2,
																																																	Left: nil,
																																																	Right: &TreeNode{
																																																		Val: 333,
																																																		Left: &TreeNode{
																																																			Val:   444,
																																																			Left:  nil,
																																																			Right: nil,
																																																		},
																																																		Right: &TreeNode{
																																																			Val:  7777777,
																																																			Left: nil,
																																																			Right: &TreeNode{
																																																				Val: 88888888,
																																																				Left: &TreeNode{
																																																					Val: 999999999,
																																																					Left: &TreeNode{
																																																						Val:  123456789,
																																																						Left: nil,
																																																						Right: &TreeNode{
																																																							Val: -777,
																																																							Left: &TreeNode{
																																																								Val:   -666,
																																																								Left:  nil,
																																																								Right: nil,
																																																							},
																																																							Right: &TreeNode{
																																																								Val: -1,
																																																								Left: &TreeNode{
																																																									Val:   0,
																																																									Left:  nil,
																																																									Right: nil,
																																																								},
																																																								Right: nil,
																																																							},
																																																						},
																																																					},
																																																					Right: nil,
																																																				},
																																																				Right: nil,
																																																			},
																																																		},
																																																	},
																																																},
																																																Right: &TreeNode{
																																																	Val:  123,
																																																	Left: nil,
																																																	Right: &TreeNode{
																																																		Val:   456,
																																																		Left:  nil,
																																																		Right: nil,
																																																	},
																																																},
																																															},
																																														},
																																													},
																																												},
																																											},
																																										},
																																									},
																																								},
																																							},
																																						},
																																					},
																																				},
																																			},
																																			Right: &TreeNode{
																																				Val: 1,
																																				Left: &TreeNode{
																																					Val:  2,
																																					Left: nil,
																																					Right: &TreeNode{
																																						Val: 1000,
																																						Left: &TreeNode{
																																							Val:   444,
																																							Left:  nil,
																																							Right: nil,
																																						},
																																						Right: &TreeNode{
																																							Val:  7777777,
																																							Left: nil,
																																							Right: &TreeNode{
																																								Val: 88888888,
																																								Left: &TreeNode{
																																									Val: 999999999,
																																									Left: &TreeNode{
																																										Val:  123456789,
																																										Left: nil,
																																										Right: &TreeNode{
																																											Val: -777,
																																											Left: &TreeNode{
																																												Val:   -666,
																																												Left:  nil,
																																												Right: nil,
																																											},
																																											Right: &TreeNode{
																																												Val: -1,
																																												Left: &TreeNode{
																																													Val:   0,
																																													Left:  nil,
																																													Right: nil,
																																												},
																																												Right: nil,
																																											},
																																										},
																																									},
																																									Right: nil,
																																								},
																																								Right: nil,
																																							},
																																						},
																																					},
																																				},
																																				Right: &TreeNode{
																																					Val:  123,
																																					Left: nil,
																																					Right: &TreeNode{
																																						Val:  456,
																																						Left: nil,
																																						Right: &TreeNode{
																																							Val: 1,
																																							Left: &TreeNode{
																																								Val:  2,
																																								Left: nil,
																																								Right: &TreeNode{
																																									Val: 333,
																																									Left: &TreeNode{
																																										Val:   444,
																																										Left:  nil,
																																										Right: nil,
																																									},
																																									Right: &TreeNode{
																																										Val:  7777777,
																																										Left: nil,
																																										Right: &TreeNode{
																																											Val: 88888888,
																																											Left: &TreeNode{
																																												Val: 999999999,
																																												Left: &TreeNode{
																																													Val:  123456789,
																																													Left: nil,
																																													Right: &TreeNode{
																																														Val: -777,
																																														Left: &TreeNode{
																																															Val:   -666,
																																															Left:  nil,
																																															Right: nil,
																																														},
																																														Right: &TreeNode{
																																															Val: -1,
																																															Left: &TreeNode{
																																																Val:   0,
																																																Left:  nil,
																																																Right: nil,
																																															},
																																															Right: nil,
																																														},
																																													},
																																												},
																																												Right: nil,
																																											},
																																											Right: nil,
																																										},
																																									},
																																								},
																																							},
																																							Right: &TreeNode{
																																								Val:  123,
																																								Left: nil,
																																								Right: &TreeNode{
																																									Val:  456,
																																									Left: nil,
																																									Right: &TreeNode{
																																										Val: 1,
																																										Left: &TreeNode{
																																											Val:  2,
																																											Left: nil,
																																											Right: &TreeNode{
																																												Val: 333,
																																												Left: &TreeNode{
																																													Val:   444,
																																													Left:  nil,
																																													Right: nil,
																																												},
																																												Right: &TreeNode{
																																													Val:  7777777,
																																													Left: nil,
																																													Right: &TreeNode{
																																														Val: 88888888,
																																														Left: &TreeNode{
																																															Val: 999999999,
																																															Left: &TreeNode{
																																																Val:  123456789,
																																																Left: nil,
																																																Right: &TreeNode{
																																																	Val: -777,
																																																	Left: &TreeNode{
																																																		Val:   -666,
																																																		Left:  nil,
																																																		Right: nil,
																																																	},
																																																	Right: &TreeNode{
																																																		Val: -1,
																																																		Left: &TreeNode{
																																																			Val:   0,
																																																			Left:  nil,
																																																			Right: nil,
																																																		},
																																																		Right: nil,
																																																	},
																																																},
																																															},
																																															Right: nil,
																																														},
																																														Right: nil,
																																													},
																																												},
																																											},
																																										},
																																										Right: &TreeNode{
																																											Val:  123,
																																											Left: nil,
																																											Right: &TreeNode{
																																												Val:  456,
																																												Left: nil,
																																												Right: &TreeNode{
																																													Val: 1,
																																													Left: &TreeNode{
																																														Val:  2,
																																														Left: nil,
																																														Right: &TreeNode{
																																															Val: 333,
																																															Left: &TreeNode{
																																																Val:   444,
																																																Left:  nil,
																																																Right: nil,
																																															},
																																															Right: &TreeNode{
																																																Val:  7777777,
																																																Left: nil,
																																																Right: &TreeNode{
																																																	Val: 88888888,
																																																	Left: &TreeNode{
																																																		Val: 999999999,
																																																		Left: &TreeNode{
																																																			Val:  123456789,
																																																			Left: nil,
																																																			Right: &TreeNode{
																																																				Val: -777,
																																																				Left: &TreeNode{
																																																					Val:   -666,
																																																					Left:  nil,
																																																					Right: nil,
																																																				},
																																																				Right: &TreeNode{
																																																					Val: -1,
																																																					Left: &TreeNode{
																																																						Val:   0,
																																																						Left:  nil,
																																																						Right: nil,
																																																					},
																																																					Right: nil,
																																																				},
																																																			},
																																																		},
																																																		Right: nil,
																																																	},
																																																	Right: nil,
																																																},
																																															},
																																														},
																																													},
																																													Right: &TreeNode{
																																														Val:  123,
																																														Left: nil,
																																														Right: &TreeNode{
																																															Val:  456,
																																															Left: nil,
																																															Right: &TreeNode{
																																																Val: 1,
																																																Left: &TreeNode{
																																																	Val:  2,
																																																	Left: nil,
																																																	Right: &TreeNode{
																																																		Val: 333,
																																																		Left: &TreeNode{
																																																			Val:   444,
																																																			Left:  nil,
																																																			Right: nil,
																																																		},
																																																		Right: &TreeNode{
																																																			Val:  7777777,
																																																			Left: nil,
																																																			Right: &TreeNode{
																																																				Val: 88888888,
																																																				Left: &TreeNode{
																																																					Val: 999999999,
																																																					Left: &TreeNode{
																																																						Val:  123456789,
																																																						Left: nil,
																																																						Right: &TreeNode{
																																																							Val: -777,
																																																							Left: &TreeNode{
																																																								Val:   -666,
																																																								Left:  nil,
																																																								Right: nil,
																																																							},
																																																							Right: &TreeNode{
																																																								Val: -1,
																																																								Left: &TreeNode{
																																																									Val:   0,
																																																									Left:  nil,
																																																									Right: nil,
																																																								},
																																																								Right: nil,
																																																							},
																																																						},
																																																					},
																																																					Right: nil,
																																																				},
																																																				Right: nil,
																																																			},
																																																		},
																																																	},
																																																},
																																																Right: &TreeNode{
																																																	Val:  123,
																																																	Left: nil,
																																																	Right: &TreeNode{
																																																		Val:   456,
																																																		Left:  nil,
																																																		Right: nil,
																																																	},
																																																},
																																															},
																																														},
																																													},
																																												},
																																											},
																																										},
																																									},
																																								},
																																							},
																																						},
																																					},
																																				},
																																			},
																																		},
																																	},
																																},
																															},
																														},
																													},
																												},
																											},
																										},
																									},
																								},
																							},
																						},
																					},
																				},
																			},
																		},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
					Right: &TreeNode{
						Val: 1000,
						Left: &TreeNode{
							Val:   444,
							Left:  nil,
							Right: nil,
						},
						Right: &TreeNode{
							Val:  7777777,
							Left: nil,
							Right: &TreeNode{
								Val: 88888888,
								Left: &TreeNode{
									Val: 999999999,
									Left: &TreeNode{
										Val:  123456789,
										Left: nil,
										Right: &TreeNode{
											Val: -777,
											Left: &TreeNode{
												Val:   -666,
												Left:  nil,
												Right: nil,
											},
											Right: &TreeNode{
												Val: -1,
												Left: &TreeNode{
													Val:   0,
													Left:  nil,
													Right: nil,
												},
												Right: nil,
											},
										},
									},
									Right: nil,
								},
								Right: nil,
							},
						},
					},
				},
				Right: &TreeNode{
					Val:  123,
					Left: nil,
					Right: &TreeNode{
						Val:  456,
						Left: nil,
						Right: &TreeNode{
							Val: 1,
							Left: &TreeNode{
								Val:  2,
								Left: nil,
								Right: &TreeNode{
									Val: 333,
									Left: &TreeNode{
										Val:   444,
										Left:  nil,
										Right: nil,
									},
									Right: &TreeNode{
										Val:  7777777,
										Left: nil,
										Right: &TreeNode{
											Val: 88888888,
											Left: &TreeNode{
												Val: 999999999,
												Left: &TreeNode{
													Val:  123456789,
													Left: nil,
													Right: &TreeNode{
														Val: -777,
														Left: &TreeNode{
															Val:   -666,
															Left:  nil,
															Right: nil,
														},
														Right: &TreeNode{
															Val: -1,
															Left: &TreeNode{
																Val:   0,
																Left:  nil,
																Right: nil,
															},
															Right: nil,
														},
													},
												},
												Right: nil,
											},
											Right: nil,
										},
									},
								},
							},
							Right: &TreeNode{
								Val:  123,
								Left: nil,
								Right: &TreeNode{
									Val:  456,
									Left: nil,
									Right: &TreeNode{
										Val: 1,
										Left: &TreeNode{
											Val:  2,
											Left: nil,
											Right: &TreeNode{
												Val: 333,
												Left: &TreeNode{
													Val:   444,
													Left:  nil,
													Right: nil,
												},
												Right: &TreeNode{
													Val:  7777777,
													Left: nil,
													Right: &TreeNode{
														Val: 88888888,
														Left: &TreeNode{
															Val: 999999999,
															Left: &TreeNode{
																Val:  123456789,
																Left: nil,
																Right: &TreeNode{
																	Val: -777,
																	Left: &TreeNode{
																		Val:   -666,
																		Left:  nil,
																		Right: nil,
																	},
																	Right: &TreeNode{
																		Val: -1,
																		Left: &TreeNode{
																			Val:   0,
																			Left:  nil,
																			Right: nil,
																		},
																		Right: nil,
																	},
																},
															},
															Right: nil,
														},
														Right: nil,
													},
												},
											},
										},
										Right: &TreeNode{
											Val:  123,
											Left: nil,
											Right: &TreeNode{
												Val:  456,
												Left: nil,
												Right: &TreeNode{
													Val: 1,
													Left: &TreeNode{
														Val:  2,
														Left: nil,
														Right: &TreeNode{
															Val: 333,
															Left: &TreeNode{
																Val:   444,
																Left:  nil,
																Right: nil,
															},
															Right: &TreeNode{
																Val:  7777777,
																Left: nil,
																Right: &TreeNode{
																	Val: 88888888,
																	Left: &TreeNode{
																		Val: 999999999,
																		Left: &TreeNode{
																			Val:  123456789,
																			Left: nil,
																			Right: &TreeNode{
																				Val: -777,
																				Left: &TreeNode{
																					Val:   -666,
																					Left:  nil,
																					Right: nil,
																				},
																				Right: &TreeNode{
																					Val: -1,
																					Left: &TreeNode{
																						Val:   0,
																						Left:  nil,
																						Right: nil,
																					},
																					Right: nil,
																				},
																			},
																		},
																		Right: nil,
																	},
																	Right: nil,
																},
															},
														},
													},
													Right: &TreeNode{
														Val:  123,
														Left: nil,
														Right: &TreeNode{
															Val:  456,
															Left: nil,
															Right: &TreeNode{
																Val: 1,
																Left: &TreeNode{
																	Val:  2,
																	Left: nil,
																	Right: &TreeNode{
																		Val: 333,
																		Left: &TreeNode{
																			Val:   444,
																			Left:  nil,
																			Right: nil,
																		},
																		Right: &TreeNode{
																			Val:  7777777,
																			Left: nil,
																			Right: &TreeNode{
																				Val: 88888888,
																				Left: &TreeNode{
																					Val: 999999999,
																					Left: &TreeNode{
																						Val:  123456789,
																						Left: nil,
																						Right: &TreeNode{
																							Val: -777,
																							Left: &TreeNode{
																								Val:   -666,
																								Left:  nil,
																								Right: nil,
																							},
																							Right: &TreeNode{
																								Val: -1,
																								Left: &TreeNode{
																									Val:   0,
																									Left:  nil,
																									Right: nil,
																								},
																								Right: nil,
																							},
																						},
																					},
																					Right: nil,
																				},
																				Right: nil,
																			},
																		},
																	},
																},
																Right: &TreeNode{
																	Val:  123,
																	Left: nil,
																	Right: &TreeNode{
																		Val:  456,
																		Left: nil,
																		Right: &TreeNode{
																			Val: 1,
																			Left: &TreeNode{
																				Val:  2,
																				Left: nil,
																				Right: &TreeNode{
																					Val: 1000,
																					Left: &TreeNode{
																						Val:   444,
																						Left:  nil,
																						Right: nil,
																					},
																					Right: &TreeNode{
																						Val:  7777777,
																						Left: nil,
																						Right: &TreeNode{
																							Val: 88888888,
																							Left: &TreeNode{
																								Val: 999999999,
																								Left: &TreeNode{
																									Val:  123456789,
																									Left: nil,
																									Right: &TreeNode{
																										Val: -777,
																										Left: &TreeNode{
																											Val:   -666,
																											Left:  nil,
																											Right: nil,
																										},
																										Right: &TreeNode{
																											Val: -1,
																											Left: &TreeNode{
																												Val:   0,
																												Left:  nil,
																												Right: nil,
																											},
																											Right: nil,
																										},
																									},
																								},
																								Right: nil,
																							},
																							Right: nil,
																						},
																					},
																				},
																			},
																			Right: &TreeNode{
																				Val:  123,
																				Left: nil,
																				Right: &TreeNode{
																					Val:  456,
																					Left: nil,
																					Right: &TreeNode{
																						Val: 1,
																						Left: &TreeNode{
																							Val:  2,
																							Left: nil,
																							Right: &TreeNode{
																								Val: 333,
																								Left: &TreeNode{
																									Val:   444,
																									Left:  nil,
																									Right: nil,
																								},
																								Right: &TreeNode{
																									Val:  7777777,
																									Left: nil,
																									Right: &TreeNode{
																										Val: 88888888,
																										Left: &TreeNode{
																											Val: 999999999,
																											Left: &TreeNode{
																												Val:  123456789,
																												Left: nil,
																												Right: &TreeNode{
																													Val: -777,
																													Left: &TreeNode{
																														Val:   -666,
																														Left:  nil,
																														Right: nil,
																													},
																													Right: &TreeNode{
																														Val: -1,
																														Left: &TreeNode{
																															Val:   0,
																															Left:  nil,
																															Right: nil,
																														},
																														Right: nil,
																													},
																												},
																											},
																											Right: nil,
																										},
																										Right: nil,
																									},
																								},
																							},
																						},
																						Right: &TreeNode{
																							Val:  123,
																							Left: nil,
																							Right: &TreeNode{
																								Val:  456,
																								Left: nil,
																								Right: &TreeNode{
																									Val: 1,
																									Left: &TreeNode{
																										Val:  2,
																										Left: nil,
																										Right: &TreeNode{
																											Val: 333,
																											Left: &TreeNode{
																												Val:   444,
																												Left:  nil,
																												Right: nil,
																											},
																											Right: &TreeNode{
																												Val:  7777777,
																												Left: nil,
																												Right: &TreeNode{
																													Val: 88888888,
																													Left: &TreeNode{
																														Val: 999999999,
																														Left: &TreeNode{
																															Val:  123456789,
																															Left: nil,
																															Right: &TreeNode{
																																Val: -777,
																																Left: &TreeNode{
																																	Val:   -666,
																																	Left:  nil,
																																	Right: nil,
																																},
																																Right: &TreeNode{
																																	Val: -1,
																																	Left: &TreeNode{
																																		Val:   0,
																																		Left:  nil,
																																		Right: nil,
																																	},
																																	Right: nil,
																																},
																															},
																														},
																														Right: nil,
																													},
																													Right: nil,
																												},
																											},
																										},
																									},
																									Right: &TreeNode{
																										Val:  123,
																										Left: nil,
																										Right: &TreeNode{
																											Val:  456,
																											Left: nil,
																											Right: &TreeNode{
																												Val: 1,
																												Left: &TreeNode{
																													Val:  2,
																													Left: nil,
																													Right: &TreeNode{
																														Val: 333,
																														Left: &TreeNode{
																															Val:   444,
																															Left:  nil,
																															Right: nil,
																														},
																														Right: &TreeNode{
																															Val:  7777777,
																															Left: nil,
																															Right: &TreeNode{
																																Val: 88888888,
																																Left: &TreeNode{
																																	Val: 999999999,
																																	Left: &TreeNode{
																																		Val:  123456789,
																																		Left: nil,
																																		Right: &TreeNode{
																																			Val: -777,
																																			Left: &TreeNode{
																																				Val:   -666,
																																				Left:  nil,
																																				Right: nil,
																																			},
																																			Right: &TreeNode{
																																				Val: -1,
																																				Left: &TreeNode{
																																					Val:   0,
																																					Left:  nil,
																																					Right: nil,
																																				},
																																				Right: nil,
																																			},
																																		},
																																	},
																																	Right: nil,
																																},
																																Right: nil,
																															},
																														},
																													},
																												},
																												Right: &TreeNode{
																													Val:  123,
																													Left: nil,
																													Right: &TreeNode{
																														Val:  456,
																														Left: nil,
																														Right: &TreeNode{
																															Val: 1,
																															Left: &TreeNode{
																																Val:  2,
																																Left: nil,
																																Right: &TreeNode{
																																	Val: 333,
																																	Left: &TreeNode{
																																		Val:   444,
																																		Left:  nil,
																																		Right: nil,
																																	},
																																	Right: &TreeNode{
																																		Val:  7777777,
																																		Left: nil,
																																		Right: &TreeNode{
																																			Val: 88888888,
																																			Left: &TreeNode{
																																				Val: 999999999,
																																				Left: &TreeNode{
																																					Val:  123456789,
																																					Left: nil,
																																					Right: &TreeNode{
																																						Val: -777,
																																						Left: &TreeNode{
																																							Val:   -666,
																																							Left:  nil,
																																							Right: nil,
																																						},
																																						Right: &TreeNode{
																																							Val: -1,
																																							Left: &TreeNode{
																																								Val:   0,
																																								Left:  nil,
																																								Right: nil,
																																							},
																																							Right: nil,
																																						},
																																					},
																																				},
																																				Right: nil,
																																			},
																																			Right: nil,
																																		},
																																	},
																																},
																															},
																															Right: &TreeNode{
																																Val:  123,
																																Left: nil,
																																Right: &TreeNode{
																																	Val: 456,
																																	Left: &TreeNode{
																																		Val: 1,
																																		Left: &TreeNode{
																																			Val:  2,
																																			Left: nil,
																																			Right: &TreeNode{
																																				Val: 1000,
																																				Left: &TreeNode{
																																					Val:   444,
																																					Left:  nil,
																																					Right: nil,
																																				},
																																				Right: &TreeNode{
																																					Val:  7777777,
																																					Left: nil,
																																					Right: &TreeNode{
																																						Val: 88888888,
																																						Left: &TreeNode{
																																							Val: 999999999,
																																							Left: &TreeNode{
																																								Val:  123456789,
																																								Left: nil,
																																								Right: &TreeNode{
																																									Val: -777,
																																									Left: &TreeNode{
																																										Val:   -666,
																																										Left:  nil,
																																										Right: nil,
																																									},
																																									Right: &TreeNode{
																																										Val: -1,
																																										Left: &TreeNode{
																																											Val:   0,
																																											Left:  nil,
																																											Right: nil,
																																										},
																																										Right: nil,
																																									},
																																								},
																																							},
																																							Right: nil,
																																						},
																																						Right: nil,
																																					},
																																				},
																																			},
																																		},
																																		Right: &TreeNode{
																																			Val:  123,
																																			Left: nil,
																																			Right: &TreeNode{
																																				Val:  456,
																																				Left: nil,
																																				Right: &TreeNode{
																																					Val: 1,
																																					Left: &TreeNode{
																																						Val:  2,
																																						Left: nil,
																																						Right: &TreeNode{
																																							Val: 333,
																																							Left: &TreeNode{
																																								Val:   444,
																																								Left:  nil,
																																								Right: nil,
																																							},
																																							Right: &TreeNode{
																																								Val:  7777777,
																																								Left: nil,
																																								Right: &TreeNode{
																																									Val: 88888888,
																																									Left: &TreeNode{
																																										Val: 999999999,
																																										Left: &TreeNode{
																																											Val:  123456789,
																																											Left: nil,
																																											Right: &TreeNode{
																																												Val: -777,
																																												Left: &TreeNode{
																																													Val:   -666,
																																													Left:  nil,
																																													Right: nil,
																																												},
																																												Right: &TreeNode{
																																													Val: -1,
																																													Left: &TreeNode{
																																														Val:   0,
																																														Left:  nil,
																																														Right: nil,
																																													},
																																													Right: nil,
																																												},
																																											},
																																										},
																																										Right: nil,
																																									},
																																									Right: nil,
																																								},
																																							},
																																						},
																																					},
																																					Right: &TreeNode{
																																						Val:  123,
																																						Left: nil,
																																						Right: &TreeNode{
																																							Val:  456,
																																							Left: nil,
																																							Right: &TreeNode{
																																								Val: 1,
																																								Left: &TreeNode{
																																									Val:  2,
																																									Left: nil,
																																									Right: &TreeNode{
																																										Val: 333,
																																										Left: &TreeNode{
																																											Val:   444,
																																											Left:  nil,
																																											Right: nil,
																																										},
																																										Right: &TreeNode{
																																											Val:  7777777,
																																											Left: nil,
																																											Right: &TreeNode{
																																												Val: 88888888,
																																												Left: &TreeNode{
																																													Val: 999999999,
																																													Left: &TreeNode{
																																														Val:  123456789,
																																														Left: nil,
																																														Right: &TreeNode{
																																															Val: -777,
																																															Left: &TreeNode{
																																																Val:   -666,
																																																Left:  nil,
																																																Right: nil,
																																															},
																																															Right: &TreeNode{
																																																Val: -1,
																																																Left: &TreeNode{
																																																	Val:   0,
																																																	Left:  nil,
																																																	Right: nil,
																																																},
																																																Right: nil,
																																															},
																																														},
																																													},
																																													Right: nil,
																																												},
																																												Right: nil,
																																											},
																																										},
																																									},
																																								},
																																								Right: &TreeNode{
																																									Val:  123,
																																									Left: nil,
																																									Right: &TreeNode{
																																										Val:  456,
																																										Left: nil,
																																										Right: &TreeNode{
																																											Val: 1,
																																											Left: &TreeNode{
																																												Val:  2,
																																												Left: nil,
																																												Right: &TreeNode{
																																													Val: 333,
																																													Left: &TreeNode{
																																														Val:   444,
																																														Left:  nil,
																																														Right: nil,
																																													},
																																													Right: &TreeNode{
																																														Val:  7777777,
																																														Left: nil,
																																														Right: &TreeNode{
																																															Val: 88888888,
																																															Left: &TreeNode{
																																																Val: 999999999,
																																																Left: &TreeNode{
																																																	Val:  123456789,
																																																	Left: nil,
																																																	Right: &TreeNode{
																																																		Val: -777,
																																																		Left: &TreeNode{
																																																			Val:   -666,
																																																			Left:  nil,
																																																			Right: nil,
																																																		},
																																																		Right: &TreeNode{
																																																			Val: -1,
																																																			Left: &TreeNode{
																																																				Val:   0,
																																																				Left:  nil,
																																																				Right: nil,
																																																			},
																																																			Right: nil,
																																																		},
																																																	},
																																																},
																																																Right: nil,
																																															},
																																															Right: nil,
																																														},
																																													},
																																												},
																																											},
																																											Right: &TreeNode{
																																												Val:  123,
																																												Left: nil,
																																												Right: &TreeNode{
																																													Val:  456,
																																													Left: nil,
																																													Right: &TreeNode{
																																														Val: 1,
																																														Left: &TreeNode{
																																															Val:  2,
																																															Left: nil,
																																															Right: &TreeNode{
																																																Val: 333,
																																																Left: &TreeNode{
																																																	Val:   444,
																																																	Left:  nil,
																																																	Right: nil,
																																																},
																																																Right: &TreeNode{
																																																	Val:  7777777,
																																																	Left: nil,
																																																	Right: &TreeNode{
																																																		Val: 88888888,
																																																		Left: &TreeNode{
																																																			Val: 999999999,
																																																			Left: &TreeNode{
																																																				Val:  123456789,
																																																				Left: nil,
																																																				Right: &TreeNode{
																																																					Val: -777,
																																																					Left: &TreeNode{
																																																						Val:   -666,
																																																						Left:  nil,
																																																						Right: nil,
																																																					},
																																																					Right: &TreeNode{
																																																						Val: -1,
																																																						Left: &TreeNode{
																																																							Val:   0,
																																																							Left:  nil,
																																																							Right: nil,
																																																						},
																																																						Right: nil,
																																																					},
																																																				},
																																																			},
																																																			Right: nil,
																																																		},
																																																		Right: nil,
																																																	},
																																																},
																																															},
																																														},
																																														Right: &TreeNode{
																																															Val:  123,
																																															Left: nil,
																																															Right: &TreeNode{
																																																Val:   456,
																																																Left:  nil,
																																																Right: nil,
																																															},
																																														},
																																													},
																																												},
																																											},
																																										},
																																									},
																																								},
																																							},
																																						},
																																					},
																																				},
																																			},
																																		},
																																	},
																																	Right: &TreeNode{
																																		Val: 1,
																																		Left: &TreeNode{
																																			Val:  2,
																																			Left: nil,
																																			Right: &TreeNode{
																																				Val: 1000,
																																				Left: &TreeNode{
																																					Val:   444,
																																					Left:  nil,
																																					Right: nil,
																																				},
																																				Right: &TreeNode{
																																					Val:  7777777,
																																					Left: nil,
																																					Right: &TreeNode{
																																						Val: 88888888,
																																						Left: &TreeNode{
																																							Val: 999999999,
																																							Left: &TreeNode{
																																								Val:  123456789,
																																								Left: nil,
																																								Right: &TreeNode{
																																									Val: -777,
																																									Left: &TreeNode{
																																										Val:   -666,
																																										Left:  nil,
																																										Right: nil,
																																									},
																																									Right: &TreeNode{
																																										Val: -1,
																																										Left: &TreeNode{
																																											Val:   0,
																																											Left:  nil,
																																											Right: nil,
																																										},
																																										Right: nil,
																																									},
																																								},
																																							},
																																							Right: nil,
																																						},
																																						Right: nil,
																																					},
																																				},
																																			},
																																		},
																																		Right: &TreeNode{
																																			Val:  123,
																																			Left: nil,
																																			Right: &TreeNode{
																																				Val:  456,
																																				Left: nil,
																																				Right: &TreeNode{
																																					Val: 1,
																																					Left: &TreeNode{
																																						Val:  2,
																																						Left: nil,
																																						Right: &TreeNode{
																																							Val: 333,
																																							Left: &TreeNode{
																																								Val:   444,
																																								Left:  nil,
																																								Right: nil,
																																							},
																																							Right: &TreeNode{
																																								Val:  7777777,
																																								Left: nil,
																																								Right: &TreeNode{
																																									Val: 88888888,
																																									Left: &TreeNode{
																																										Val: 999999999,
																																										Left: &TreeNode{
																																											Val:  123456789,
																																											Left: nil,
																																											Right: &TreeNode{
																																												Val: -777,
																																												Left: &TreeNode{
																																													Val:   -666,
																																													Left:  nil,
																																													Right: nil,
																																												},
																																												Right: &TreeNode{
																																													Val: -1,
																																													Left: &TreeNode{
																																														Val:   0,
																																														Left:  nil,
																																														Right: nil,
																																													},
																																													Right: nil,
																																												},
																																											},
																																										},
																																										Right: nil,
																																									},
																																									Right: nil,
																																								},
																																							},
																																						},
																																					},
																																					Right: &TreeNode{
																																						Val:  123,
																																						Left: nil,
																																						Right: &TreeNode{
																																							Val:  456,
																																							Left: nil,
																																							Right: &TreeNode{
																																								Val: 1,
																																								Left: &TreeNode{
																																									Val:  2,
																																									Left: nil,
																																									Right: &TreeNode{
																																										Val: 333,
																																										Left: &TreeNode{
																																											Val:   444,
																																											Left:  nil,
																																											Right: nil,
																																										},
																																										Right: &TreeNode{
																																											Val:  7777777,
																																											Left: nil,
																																											Right: &TreeNode{
																																												Val: 88888888,
																																												Left: &TreeNode{
																																													Val: 999999999,
																																													Left: &TreeNode{
																																														Val:  123456789,
																																														Left: nil,
																																														Right: &TreeNode{
																																															Val: -777,
																																															Left: &TreeNode{
																																																Val:   -666,
																																																Left:  nil,
																																																Right: nil,
																																															},
																																															Right: &TreeNode{
																																																Val: -1,
																																																Left: &TreeNode{
																																																	Val:   0,
																																																	Left:  nil,
																																																	Right: nil,
																																																},
																																																Right: nil,
																																															},
																																														},
																																													},
																																													Right: nil,
																																												},
																																												Right: nil,
																																											},
																																										},
																																									},
																																								},
																																								Right: &TreeNode{
																																									Val:  123,
																																									Left: nil,
																																									Right: &TreeNode{
																																										Val:  456,
																																										Left: nil,
																																										Right: &TreeNode{
																																											Val: 1,
																																											Left: &TreeNode{
																																												Val:  2,
																																												Left: nil,
																																												Right: &TreeNode{
																																													Val: 333,
																																													Left: &TreeNode{
																																														Val:   444,
																																														Left:  nil,
																																														Right: nil,
																																													},
																																													Right: &TreeNode{
																																														Val:  7777777,
																																														Left: nil,
																																														Right: &TreeNode{
																																															Val: 88888888,
																																															Left: &TreeNode{
																																																Val: 999999999,
																																																Left: &TreeNode{
																																																	Val:  123456789,
																																																	Left: nil,
																																																	Right: &TreeNode{
																																																		Val: -777,
																																																		Left: &TreeNode{
																																																			Val:   -666,
																																																			Left:  nil,
																																																			Right: nil,
																																																		},
																																																		Right: &TreeNode{
																																																			Val: -1,
																																																			Left: &TreeNode{
																																																				Val:   0,
																																																				Left:  nil,
																																																				Right: nil,
																																																			},
																																																			Right: nil,
																																																		},
																																																	},
																																																},
																																																Right: nil,
																																															},
																																															Right: nil,
																																														},
																																													},
																																												},
																																											},
																																											Right: &TreeNode{
																																												Val:  123,
																																												Left: nil,
																																												Right: &TreeNode{
																																													Val:  456,
																																													Left: nil,
																																													Right: &TreeNode{
																																														Val: 1,
																																														Left: &TreeNode{
																																															Val:  2,
																																															Left: nil,
																																															Right: &TreeNode{
																																																Val: 333,
																																																Left: &TreeNode{
																																																	Val:   444,
																																																	Left:  nil,
																																																	Right: nil,
																																																},
																																																Right: &TreeNode{
																																																	Val:  7777777,
																																																	Left: nil,
																																																	Right: &TreeNode{
																																																		Val: 88888888,
																																																		Left: &TreeNode{
																																																			Val: 999999999,
																																																			Left: &TreeNode{
																																																				Val:  123456789,
																																																				Left: nil,
																																																				Right: &TreeNode{
																																																					Val: -777,
																																																					Left: &TreeNode{
																																																						Val:   -666,
																																																						Left:  nil,
																																																						Right: nil,
																																																					},
																																																					Right: &TreeNode{
																																																						Val: -1,
																																																						Left: &TreeNode{
																																																							Val:   0,
																																																							Left:  nil,
																																																							Right: nil,
																																																						},
																																																						Right: nil,
																																																					},
																																																				},
																																																			},
																																																			Right: nil,
																																																		},
																																																		Right: nil,
																																																	},
																																																},
																																															},
																																														},
																																														Right: &TreeNode{
																																															Val:  123,
																																															Left: nil,
																																															Right: &TreeNode{
																																																Val:   456,
																																																Left:  nil,
																																																Right: nil,
																																															},
																																														},
																																													},
																																												},
																																											},
																																										},
																																									},
																																								},
																																							},
																																						},
																																					},
																																				},
																																			},
																																		},
																																	},
																																},
																															},
																														},
																													},
																												},
																											},
																										},
																									},
																								},
																							},
																						},
																					},
																				},
																			},
																		},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			c := Constructor()

			ser := c.serialize(tt.root)
			des := c.deserialize(ser)

			if !reflect.DeepEqual(des, tt.root) {
				t.Fatalf("got %v, want %v", des, tt.root)
			}
		})
	}
}
